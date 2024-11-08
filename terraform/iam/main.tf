terraform {
    required_version = "~> v1.6.0" 

    required_providers {
        google = {
            source = "hashicorp/google"
            version = "6.8.0"
        }
    }

    backend "gcs" {
    bucket = "farra"
    prefix = "learn-with-farra-iam.vpc" 
    }
}

locals {
  project = "learn-with-farra"
  region = "asia-southeast2"
}

provider "google" {
  project = local.project
  region = local.region
}


# data "terraform_remote_state" "sa_farra" { #jika data ada di direktori berbeda bisa gunakan remote state untuk ngelink datanya
#   backend = "gcs"
#   config = {
#     bucket = "farra"
#     prefix = "learn-with-farra-0.vpc" #prefix yang ada di service account
#   }
# }

#iam member
# resource "google_project_iam_member" "farra-iam" {
#   for_each = toset([
#     "roles/owner",
#     "roles/editor",
#   ])

#   project = "learn-with-farra"
#   role    = each.value
#   member  = "serviceAccount:${data.terraform_remote_state.sa_farra.outputs.email_farra_sa}" #mengambil data dari terraform state dan output di tempat lain (sa)
# }

resource "google_project_iam_member" "iam1" {
  project = local.project
  role    = "roles/artifactregistry.admin"
  member  = "serviceAccount:farra-service-acc@learn-with-farra.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "iam2" {
  project = local.project
  role    = "roles/container.admin"
  member  = "serviceAccount:farra-service-acc@learn-with-farra.iam.gserviceaccount.com"
}