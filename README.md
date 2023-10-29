# Seopilot

This is the **Seopilot** project.

## Requirements

Here are the requirements:

- mkcert
- docker

```bash
# Edit your local /etc/hosts file
# Add the following lines
sudo echo "127.0.0.1 dashboard.seopilot.dev" >> /etc/hosts
sudo echo "127.0.0.1 api.seopilot.dev" >> /etc/hosts

# Install certificates for enable SSL via traefik
mkcert -install

# Move into certs folder
cd certs

# Create the local certificates
mkcert seopilot.dev "*.seopilot.dev" 127.0.0.1

# Rename to generic naming
# The pem files will be mounted as volume for traefik
mv seopilot*key.pem app-key.pem
mv seopilot*.pem app.pem

```

## Develop

```bash
docker compose build
docker compose up
```
