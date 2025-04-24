output "client_url" {
  description = "URL del balanceador de carga del cliente"
  value       = "http://${aws_lb.client_lb.dns_name}"
}

output "api_url" {
  description = "URL del balanceador de carga de la API"
  value       = "http://${aws_lb.api_lb.dns_name}"
}

output "redis_endpoint" {
  description = "Endpoint del cluster de Redis"
  value       = aws_elasticache_cluster.redis.cache_nodes.0.address
  sensitive   = true
}

output "ecs_cluster_name" {
  description = "Nombre del cluster ECS"
  value       = aws_ecs_cluster.stock_app_cluster.name
}

output "vpc_id" {
  description = "ID de la VPC"
  value       = aws_vpc.stock_app_vpc.id
}

output "subnets" {
  description = "IDs de las subredes"
  value = {
    public_subnet_1 = aws_subnet.public_subnet_1.id,
    public_subnet_2 = aws_subnet.public_subnet_2.id
  }
} 