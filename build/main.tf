locals {
  prismUser     = "husain@lab.demo"
  prismPassword = "nutanix/4u"
  prismEndpoint = "10.38.14.9"
  cluster_name  = "PHX-SPOC014-1"
  network_name  = "default"
  vm_image      = "centos-8"
  vmname_prefix = "husain"
  amount        = 2
  password      = "nutanix/4u"
}

provider "nutanix" {
  username = local.prismUser
  password = local.prismPassword
  endpoint = local.prismEndpoint
  insecure = true
  port     = "9440"
}

data "nutanix_image" "image" {
  image_name = local.vm_image
}
//
data "nutanix_projects" "projects" {}
//
//data "nutanix_clusters" "clusters" {}
//
//data "nutanix_subnets" "subnets" {}

data "nutanix_project" "project" {
  project_id = "447a90f8-87f0-484f-803f-336dc60206a8"
}

//data "nutanix_market_items" "items" {}

//output "project" {
//  value = data.nutanix_project.project
//}

//resource "nutanix_application" "application" {
//  marketitem_uuid = "9fc0303a-909d-4e0f-a5e7-9bca300eb58e"
//  project_uuid = "447a90f8-87f0-484f-803f-336dc60206a8"
//  environment_uuid = "19b7122e-b470-8e3c-9b3c-0d6a0f8ea7b9"
//  name = "somewhere-app3"
//  description = "somere first app"
//}

//output "items" {
//  value = data.nutanix_market_items.items.*
//}
//
output "projects" {
  value = data.nutanix_projects.projects.*
}

output "image" {
  value = data.nutanix_image.image
}

output "project" {
  value = data.nutanix_project.project
}