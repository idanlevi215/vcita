provider "google" {
  credentials = file("./sa.json")
  project     = "silver-aurora-311809"
  region      = "us-central1"
  version     = "~> 2.5.0"
}
