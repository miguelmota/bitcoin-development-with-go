all: build

.PHONY: install
install:
	npm install gitbook-cli@latest -g

.PHONY: serve
serve:
	gitbook serve

.PHONY: build
build:
	gitbook build

.PHONY: deploy
deploy:
	./deploy.sh

.PHONY: deploy/all
deploy/all: build pdf ebook mobi deploy

.PHONY: pdf
pdf:
	gitbook pdf ./ bitcoin-development-with-go.pdf

.PHONY: ebook
ebook:
	gitbook epub ./ bitcoin-development-with-go.epub

.PHONY: mobi
mobi:
	gitbook mobi ./ bitcoin-development-with-go.mobi

.PHONY: plugins/install
plugins/install:
	gitbook install

.PHONY: btc/hexpriv2hexpub
btc/hexpriv2hexpub:
	@bitcoin-tool --network bitcoin --input-type private-key --input-format hex --output-type public-key --public-key-compression uncompressed --output-format hex --input $(KEY)

.PHONY: btc/b58priv2b58pub
btc/b58priv2b58pub:
	@bitcoin-tool --network bitcoin --input-type private-key-wif --input-format base58check --output-type public-key --public-key-compression auto --output-format base58check --input $(KEY)

.PHONY: btc/b58priv2hexpub
btc/b58priv2hexpub:
	@bitcoin-tool --network bitcoin --input-type private-key-wif --input-format base58check --output-type public-key --public-key-compression auto --output-format hex --input $(KEY)

.PHONY: btc/hexpriv2b58address
btc/hexpriv2b58address:
	@bitcoin-tool --network bitcoin --input-type private-key --input-format hex --output-type address --public-key-compression uncompressed --output-format base58check --input $(KEY)

.PHONY: btc/b58priv2b58address
btc/b58priv2b58address:
	@bitcoin-tool --network bitcoin --input-type private-key-wif --input-format base58check --output-type address --public-key-compression auto --output-format base58check --input $(KEY)

# 15vFKWjdm2LWWoJ2B2c1Kx9tu79eANoGLV
# L2673vhaBt1UgmQAFoAV3qA2N4gJFVAPL2LdAr6FoTzjyYEizhGo
