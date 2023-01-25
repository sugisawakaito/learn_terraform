# help: ## ヘルプを表示する
#     @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: ## terraform init を実行
	terraform init

validate: ## tf ファイルを検証する
	terraform validate && printf "\e[32m***** validated ******\e[m\n"

plan: validate ## terraform plan を実行する
	terraform plan

apply: plan ## terraform apply を実行する
	terraform apply --auto-approve

show: ## terraform show を実行する
	terraform show

destroy: ## terraform destroy を実行する
	terraform destroy

refresh: ## terraform refresh を実行する
	terraform refresh:w

fmt: ## terraform fmt を実行する
	terraform fmt

# test: ## ローカル環境で circleci を実行する
#     circleci build .circleci/config.yml