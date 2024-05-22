package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/hashicorp/go-multierror"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
	"github.com/tomnomnom/linkheader"
)

type SudoCommandsBundle struct {
	Id                 *string             `json:"id"`
	Name               *string             `json:"name"`
	StructuredCommands []StructuredCommand `json:"structured_commands"`
	RunAs              *string             `json:"run_as"`
	NoPasswd           *bool               `json:"nopasswd"`
	NoExec             *bool               `json:"noexec"`
	SetEnv             *bool               `json:"setenv"`
	AddEnv             []string            `json:"add_env"`
	SubEnv             []string            `json:"sub_env"`
}
type StructuredCommand struct {
	CommandType     *string `json:"command_type"`
	Command         *string `json:"command"`
	ArgsType        *string `json:"args_type,omitempty"`
	Args            *string `json:"args,omitempty"`
	RenderedCommand *string `json:"rendered_command,omitempty"`
}

type SudoCommandsBundleListResponse struct {
	SudoCommandsBundle []SudoCommandsBundle `json:"list"`
}

func (s SudoCommandsBundle) Exists() bool {
	return utils.IsNonEmpty(s.Id)
}

func (s SudoCommandsBundle) ToResourceMap() map[string]any {
	m := make(map[string]any)

	if s.Id != nil {
		m[attributes.ID] = *s.Id
	}
	if s.Name != nil {
		m[attributes.Name] = *s.Name
	}

	structuredCommands := make([]map[string]any, len(s.StructuredCommands))
	for i, structuredCommand := range s.StructuredCommands {
		serializedStructuredCommand := make(map[string]any)
		if structuredCommand.CommandType != nil {
			serializedStructuredCommand[attributes.StructuredCommandType] = structuredCommand.CommandType
		}

		if structuredCommand.Command != nil {
			serializedStructuredCommand[attributes.StructuredCommand] = structuredCommand.Command
		}

		if structuredCommand.ArgsType != nil {
			serializedStructuredCommand[attributes.StructuredCommandArgsType] = structuredCommand.ArgsType
		}

		if structuredCommand.Args != nil {
			serializedStructuredCommand[attributes.StructuredCommandArgs] = structuredCommand.Args
		}

		if structuredCommand.RenderedCommand != nil {
			serializedStructuredCommand[attributes.StructuredRenderedCommand] = structuredCommand.RenderedCommand
		}

		structuredCommands[i] = serializedStructuredCommand
	}

	m[attributes.StructuredCommands] = structuredCommands
	return m
}

func validateAddSubEnv(name string, addSubEnvs []string, errs *multierror.Error) bool {
	addSubEnvRegex := regexp.MustCompile(`^[\w]+$`)
	hasErr := false

	for _, addSubEnv := range addSubEnvs {
		if addSubEnvRegex.MatchString(addSubEnv) {
			hasErr = true
			multierror.Append(errs, fmt.Errorf("%s is not valid", name))
		}
	}

	return hasErr
}

func isSudoCommandsBundleValid(sudoCommandsBundle SudoCommandsBundle) (bool, error) {
	var errs *multierror.Error

	if sudoCommandsBundle.Name == nil {
		multierror.Append(errs, fmt.Errorf("sudo commands bundle name is empty"))
		return false, errs
	}

	namePattern := regexp.MustCompile(`^[\w\-_.]+$`)
	runAsPattern := regexp.MustCompile(`^([%]{0,1})((([#])(\d+))|([\w\-_.]+)|((?i)[A-Z0-9._%+-^]+@[A-Z0-9.-]+\.[A-Z]{2,}))$`)

	sudoCommandTypeRawValidator := regexp.MustCompile(`^[[:print:]]+$`)
	sudoCommandTypeExecutableValidator := regexp.MustCompile(`^/([^/]+/)*[^/]+$`)
	sudoCommandTypeDirectoryValidator := regexp.MustCompile(`^(/|/([^/]+/)+)$`)
	noMatchValidator := regexp.MustCompile(`^\b$`)
	sudoCommandBundleExecutableArgsValidator := regexp.MustCompile(`^[[:print:]]*$`)
	sudoCommandBundleExecutableArgsTypeValidator := regexp.MustCompile(`^(any|none|custom)$`)

	nameValidation := len(*sudoCommandsBundle.Name) > 1 && len(*sudoCommandsBundle.Name) <= 255 && namePattern.MatchString(*sudoCommandsBundle.Name)
	if !nameValidation {
		multierror.Append(errs, fmt.Errorf("name is not valid"))
	}

	runAsValidation := sudoCommandsBundle.RunAs != nil && (len(*sudoCommandsBundle.RunAs) == 0 || runAsPattern.MatchString(*sudoCommandsBundle.RunAs))
	if !runAsValidation {
		multierror.Append(errs, fmt.Errorf("run_as is not valid"))
	}

	addEnvValidation := validateAddSubEnv("add_env", sudoCommandsBundle.AddEnv, errs)
	subEnvValidation := validateAddSubEnv("sub_env", sudoCommandsBundle.SubEnv, errs)

	commandTypeValidators := map[string]*regexp.Regexp{
		"raw":        sudoCommandTypeRawValidator,
		"directory":  sudoCommandTypeDirectoryValidator,
		"executable": sudoCommandTypeExecutableValidator,
	}

	argsValidators := map[string]*regexp.Regexp{
		"raw":        noMatchValidator,
		"directory":  noMatchValidator,
		"executable": sudoCommandBundleExecutableArgsValidator,
	}

	argsTypeValidators := map[string]*regexp.Regexp{
		"raw":        noMatchValidator,
		"directory":  noMatchValidator,
		"executable": sudoCommandBundleExecutableArgsTypeValidator,
	}

	var structuredCommandsValidation = true
	for _, structuredCommand := range sudoCommandsBundle.StructuredCommands {
		if structuredCommand.CommandType == nil {
			structuredCommandsValidation = false
			multierror.Append(errs, fmt.Errorf("command_type is not valid"))
			continue
		}

		commandValidator, ok := commandTypeValidators[*structuredCommand.CommandType]
		if !ok {
			structuredCommandsValidation = false
			multierror.Append(errs, fmt.Errorf("command_type is not valid"))
		} else {
			commandValidation := structuredCommand.Command != nil && len(*structuredCommand.Command) > 0 && commandValidator.MatchString(*structuredCommand.Command)
			if !commandValidation {
				structuredCommandsValidation = false
				multierror.Append(errs, fmt.Errorf("command is not valid"))
			}
		}

		argsValidator, ok := argsValidators[*structuredCommand.CommandType]
		if !ok {
			structuredCommandsValidation = false
			multierror.Append(errs, fmt.Errorf("command_type is not valid"))
		} else {
			argsValidation := structuredCommand.Args != nil && len(*structuredCommand.Args) > 0 && argsValidator.MatchString(*structuredCommand.Args)
			if !argsValidation {
				structuredCommandsValidation = false
				multierror.Append(errs, fmt.Errorf("args is not valid"))
			}
		}

		argsTypeValidator, ok := argsTypeValidators[*structuredCommand.CommandType]
		if !ok {
			structuredCommandsValidation = false
			multierror.Append(errs, fmt.Errorf("args_type is not valid"))
		} else {
			argsTypeValidation := structuredCommand.ArgsType != nil && len(*structuredCommand.ArgsType) > 0 && argsTypeValidator.MatchString(*structuredCommand.ArgsType)
			if !argsTypeValidation {
				structuredCommandsValidation = false
				multierror.Append(errs, fmt.Errorf("args_type is not valid"))
			}
		}
	}

	return nameValidation && runAsValidation && addEnvValidation && subEnvValidation && structuredCommandsValidation, errs
}

func (c OktaPAMClient) ListSudoCommandsBundles(ctx context.Context) ([]SudoCommandsBundle, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/sudo_command_bundles", url.PathEscape(c.Team))
	sudoCommandsBundles := make([]SudoCommandsBundle, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again.
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetResult(&SudoCommandsBundleListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		sudoCommandsBundleListResponse := resp.Result().(*SudoCommandsBundleListResponse)
		sudoCommandsBundles = append(sudoCommandsBundles, sudoCommandsBundleListResponse.SudoCommandsBundle...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""

		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
				break
			}
		}
	}

	return sudoCommandsBundles, nil
}

func (c OktaPAMClient) GetSudoCommandsBundle(ctx context.Context, sudoCommandsBundleId string) (*SudoCommandsBundle, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/sudo_command_bundles/%s", url.PathEscape(c.Team), url.PathEscape(sudoCommandsBundleId))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&CloudConnection{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		sudoCommandsBundle := resp.Result().(*SudoCommandsBundle)
		if sudoCommandsBundle.Exists() {
			return sudoCommandsBundle, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateSudoCommandsBundle(ctx context.Context, sudoCommandsBundle SudoCommandsBundle) (*SudoCommandsBundle, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/sudo_command_bundles", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)

	if valid, errs := isSudoCommandsBundleValid(sudoCommandsBundle); !valid {
		fmt.Println("Error validating sudo commands bundle data", errs)
		return nil, errs
	}

	resp, err := c.CreateBaseRequest(ctx).SetBody(sudoCommandsBundle).SetResult(&SudoCommandsBundle{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err = checkStatusCode(resp, http.StatusCreated); err != nil {
		return nil, err
	}
	return resp.Result().(*SudoCommandsBundle), nil
}

func (c OktaPAMClient) UpdateSudoCommandsBundle(ctx context.Context, sudoCommandsBundleId string, updates map[string]any) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/sudo_command_bundles/%s", url.PathEscape(c.Team), url.PathEscape(sudoCommandsBundleId))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err = checkStatusCode(resp, http.StatusNoContent); err != nil {
		return err
	}
	return nil
}

func (c OktaPAMClient) DeleteSudoCommandsBundle(ctx context.Context, sudoCommandsBundleId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/sudo_command_bundles/%s", url.PathEscape(c.Team), url.PathEscape(sudoCommandsBundleId))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	if _, err = checkStatusCode(resp, http.StatusOK, http.StatusNoContent, http.StatusNotFound); err != nil {
		return err
	}

	return nil
}
