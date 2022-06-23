resource "oktapam_kubernetes_cluster" "k3s" {
  key  = "k3s"
  auth = OIDC_RSA2048

  labels = {
    foo = "bar"
    baz = "qux"
  }

  // outputs:
  // id
  // oidc_issuer_url
  // labels (? now with clusterkey= ... in it?)
}

resource "cluster_goes_here" "k3s" {
  make_me_a_cluster = {
    oidc_issuer_url = oktapam_kubernetes_cluster.k3s.oidc_issuer_url
  }

  // output
  // public_certificate
  // url <- control plane url
}

resource "oktapam_kubernetes_cluster_connection" "k3s" {
  cluster_id         = oktapam_kubernetes_cluster.k3s.id
  public_certificate = cluster_goes_here.k3s.public_certificate
  api_url            =  cluster_goes_here.k3s.url
}

resource "oktapam_kubernetes_cluster_group" "everyone" {
  group_name = oktapam_project_group.everyone.group_name
  selector   = "clusterkey=${oktapam_kubernetes_cluster.k3s.key}"
  claims     = {
    groups = ["foo", "bar"]
  }
}
