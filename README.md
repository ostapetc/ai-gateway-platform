### Generate service code by service.api file
goctl api go -api service.api -dir . 

### Update all submodules at once
git submodule update --remote --recursive