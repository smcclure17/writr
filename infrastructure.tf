resource "google_redis_instance" "redis-cache" {
  name           = "document-memory-cache"
  project        = "writr-dev-384017"
  tier           = "STANDARD_HA"
  memory_size_gb = 1

  location_id             = "us-central1-a"
  alternative_location_id = "us-central1-f"

  # authorized_network = data.google_compute_network.redis-network.id

  redis_version = "REDIS_4_0"
  display_name  = "Document Memory Cache"
  region        = "us-central1"
}


resource "google_project_service" "vpcaccess-api" {
  project = "writr-dev-384017"
  service = "vpcaccess.googleapis.com"
}

module "test-vpc-module" {
  source       = "terraform-google-modules/network/google"
  version      = "~> 6.0"
  project_id   = "writr-dev-384017"
  network_name = "redis-network"
  mtu          = 1460

  subnets = [
    {
      subnet_name   = "serverless-subnet"
      subnet_ip     = "10.10.10.0/28"
      subnet_region = "us-central1"
    }
  ]
}

# TODO Add datastore config
