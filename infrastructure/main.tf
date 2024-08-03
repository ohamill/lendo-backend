terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "5.39.1"
    }
  }
}

provider "google" {
  project = "owen-playground"
}

resource "google_artifact_registry_repository" "lendo" {
  location      = "europe-west1"
  repository_id = "lendo"
  format        = "DOCKER"
}