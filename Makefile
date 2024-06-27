NUM := $(shell echo "${PROJECT}" | cut -d '-' -f1)

new:
	mkdir -p "$(PROJECT)"
	go mod init baris/lc/$(NUM)
	mv go.mod "$(PROJECT)"/
	touch main.go
	mv main.go "$(PROJECT)"/
	cp Makefile.tmp "$(PROJECT)/Makefile"
	cd "$(PROJECT)" && code ./
