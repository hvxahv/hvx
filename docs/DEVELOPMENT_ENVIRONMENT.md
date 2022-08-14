# DEVELOPMENT ENVIRONMENT DEPLOY

Infrastructure build for development environments, using terraform.
https://github.com/hvxahv/infra4dev

---
.tf: # Language: terraform

Using the local configuration file: [README](../conf/README.md)

Start the service：

```bash
cd hack/
```
Go to the following folder: [hack/](../hack)

Perform the following services:
- ./run.sh gw
- ./run.sh public
- ./run.sh account
- ./run.sh actor
- ./run.sh auth
- ./run.sh device
- ./run.sh article
- ./run.sh channel
- ./run.sh saved
- ./run.sh activity
- ./run.sh message
- ./run.sh notify