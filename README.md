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

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/master/Zadanie%201%20-%20Docker/demo)



**Zadanie 2 - Scala**

Należy stworzyć aplikację na frameworku Play w Scali 2.

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów

:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy

:white_check_mark: 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD

:white_check_mark: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok

:white_check_mark: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all, show by id (get), update (put), delete (delete), add (post).

[Link do commita](https://github.com/piotrklosowski96/ebiznes24/commit/e8aea0089a062fcf5bb83e045af60f784c23ab40)

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/master/Zadanie%202%20-%20Scala/demo)



**Zadanie 3 - Kotlin**

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord

[Link do commita 1](https://github.com/piotrklosowski96/ebiznes24/tree/ae3e95cd7bff8bf824224ada347b649455e25048)

:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota)

[Link do commita 2](https://github.com/piotrklosowski96/ebiznes24/tree/22f648974e2257afdb7f7fab183fd63a0d95eec7)

:white_check_mark: 4.0 Zwróci listę kategorii na określone żądanie użytkownika

[Link do commita 3](https://github.com/piotrklosowski96/ebiznes24/tree/98834633f8ba694cbcacb5adcbf157dcc5b334b1)

:white_check_mark: 4.5 Zwróci listę produktów wg żądanej kategorii

[Link do commita 4](https://github.com/piotrklosowski96/ebiznes24/tree/98834633f8ba694cbcacb5adcbf157dcc5b334b1)

:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenfer, Webex, Skype, Discrod

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/57d010221659d4e1765f5dd414734715ca02e99d/Zadanie%203%20-%20Kotlin/demo)



**Zadanie 4 - Go**

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD

[Link do commita 1](https://github.com/piotrklosowski96/ebiznes24/tree/130b2e018c8ea73e18f99667d79701aa70778eaa)

[Link do commita 2](https://github.com/piotrklosowski96/ebiznes24/tree/fae3bce2e73016079773a668544b10975db1b6b8)

[Link do commita 3](https://github.com/piotrklosowski96/ebiznes24/tree/e987095ac632c630a012e9987bb1f6eee2699537)

:white_check_mark: 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)

[Link do commita 4](https://github.com/piotrklosowski96/ebiznes24/tree/e1b12ad438a2b19d8c405d1753c403720f054f1c)

:white_check_mark: 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint

[Link do commita 5](https://github.com/piotrklosowski96/ebiznes24/tree/2b7e6bbc5857d983986f504b7e69573bb4d638b2)

:white_check_mark: 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem

[Link do commita 6](https://github.com/piotrklosowski96/ebiznes24/tree/e7f2e8b48c640c105dd41a37357d33d28bb42cfd)

:white_check_mark: 5.0 pogrupować zapytania w gorm’owe scope'y

[Link do commita 7](https://github.com/piotrklosowski96/ebiznes24/tree/185d7f0979d6a6d41452881e05f54d412aa528eb)

Poprawki zaimplementowałem z małym (~1h) poślizgiem - demo uwzględnia je:

[Link do poprawki 1](https://github.com/piotrklosowski96/ebiznes24/tree/0bac0f2d200f28ebe8d646751638c7a99746f973)

[Link do poprawki 2](https://github.com/piotrklosowski96/ebiznes24/tree/92bdb5c28eb894ad9118d7b7de6cc654097ceb66)

[Link do poprawki 3](https://github.com/piotrklosowski96/ebiznes24/tree/c9e272e2845f52cb24d74fd24ba46149a1e8f0ba)

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/57d010221659d4e1765f5dd414734715ca02e99d/Zadanie%204%20-%20Go/demo)