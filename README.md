**Zadanie 1 - Docker** 

Należy stworzyć obraz oraz umieścić obraz na hub.docker.com, a Dockerfile na githubie.

:white_check_mark: 3.0 Obraz ubuntu z Pythonem w wersji 3.8

[Link do commita 1](https://github.com/piotrklosowski96/ebiznes24/commit/2b35f165f9da2ce5ddc699f3bbe7a52da2ab83f0)

:white_check_mark: 3.5 Obraz ubuntu:22.04 z Javą w wersji 8 oraz Kotlinem

[Link do commita 2](https://github.com/piotrklosowski96/ebiznes24/commit/113587773f0732bb8b2552ba20e6b8300f89374f)

:white_check_mark: 4.0 Do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle)

[Link do commita 3](https://github.com/piotrklosowski96/ebiznes24/commit/27f71e4fe35166da59bc861eafee10251e0eef23)

:white_check_mark: 4.5 Należy stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle

[Link do commita 4](https://github.com/piotrklosowski96/ebiznes24/commit/27f71e4fe35166da59bc861eafee10251e0eef23)

:white_check_mark: 5.0 Należy dodać konfigurację docker-compose

[Link do commita 5](https://github.com/piotrklosowski96/ebiznes24/commit/f6a73d980895a7ed98a23d1b793eee4a8a591748)

Wymaganie 3 i 4 zaimplementowałem w tym samym commicie.

Kod: [Folder z zadaniem 1](https://github.com/piotrklosowski96/ebiznes24/tree/master/Zadanie%201%20-%20Docker)

Link do obrazu na hubie: [Link do hub.docker.com](https://hub.docker.com/layers/piotrklosowski/ebiznes24/latest/images/sha256-8bb50a142475e4689a0c86ad12058ca7a402c434abe0680c84d5aa485d871e13?context=repo)



**Zadanie 2 - Scala**

Należy stworzyć aplikację na frameworku Play w Scali 2.

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów

:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy

:white_check_mark: 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD

:white_check_mark: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok

:white_check_mark: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all, show by id (get), update (put), delete (delete), add (post).

[Link do commita](https://github.com/piotrklosowski96/ebiznes24/commit/e8aea0089a062fcf5bb83e045af60f784c23ab40)