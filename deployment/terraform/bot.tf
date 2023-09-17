# resource "kubernetes_deployment" "bot" {
#   metadata {
#     name = var.gke_deployment_name
#     labels = {
#       app = var.gke_deployment_name
#     }
#   }

#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = var.gke_deployment_name
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = var.gke_deployment_name
#         }
#       }
#       spec {
#         container {
#           name    = var.gke_deployment_name
#           image   = var.gcr_image_name
#           command = ["./hayasui"]
#           args    = ["bot"]
#           env {
#             name  = "HYS_DISCORD_TOKEN"
#             value = var.hys_discord_token
#           }
#           env {
#             name  = "HYS_DISCORD_PREFIX"
#             value = var.hys_discord_prefix
#           }
#           env {
#             name  = "HYS_CACHE_DIALECT"
#             value = var.hys_cache_dialect
#           }
#           env {
#             name  = "HYS_CACHE_ADDRESS"
#             value = var.hys_cache_address
#           }
#           env {
#             name  = "HYS_CACHE_PASSWORD"
#             value = var.hys_cache_password
#           }
#           env {
#             name  = "HYS_CACHE_TIME"
#             value = var.hys_cache_time
#           }
#           env {
#             name  = "HYS_LOG_LEVEL"
#             value = var.hys_log_level
#           }
#           env {
#             name  = "HYS_LOG_JSON"
#             value = var.hys_log_json
#           }
#           env {
#             name  = "HYS_NEWRELIC_LICENSE_KEY"
#             value = var.hys_newrelic_license_key
#           }
#         }
#       }
#     }
#   }
# }
