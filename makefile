include .env

show_env: ## 環境変数を表示する
	@echo $(TARGET_EC2_IMAGE)
	@echo $(KEY_PAIR_NAME)

help: ## ヘルプを表示する
    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: ## terraform init を実行
	terraform init

validate: ## tf ファイルを検証する
	terraform validate && printf "\e[32m***** validated ******\e[m\n"

plan: fmt ## terraform plan を実行する
	terraform plan

apply: plan ## terraform apply を実行する
	terraform apply --auto-approve

destroy: ## terraform destroy を実行する
	terraform destroy

fmt: validate ## terraform fmt を実行する
	terraform fmt

# test: ## ローカル環境で circleci を実行する
#     circleci build .circleci/config.yml

ec2_describe_images:
	aws ec2 describe-images --image-ids $(TARGET_EC2_IMAGE)

ssh-keygen:
	scripts/make_keypair.sh