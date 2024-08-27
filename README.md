**Zadanie 1 - Docker** 

Należy stworzyć obraz oraz umieścić obraz na hub.docker.com, a Dockerfile na githubie.

:white_check_mark: 3.0 Obraz ubuntu z Pythonem w wersji 3.8<br>
[Link do commita 1](https://github.com/piotrklosowski96/ebiznes24/commit/2b35f165f9da2ce5ddc699f3bbe7a52da2ab83f0)

:white_check_mark: 3.5 Obraz ubuntu:22.04 z Javą w wersji 8 oraz Kotlinem<br>
[Link do commita 2](https://github.com/piotrklosowski96/ebiznes24/commit/113587773f0732bb8b2552ba20e6b8300f89374f)

:white_check_mark: 4.0 Do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle)<br>
[Link do commita 3](https://github.com/piotrklosowski96/ebiznes24/commit/27f71e4fe35166da59bc861eafee10251e0eef23)

:white_check_mark: 4.5 Należy stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle<br>
[Link do commita 4](https://github.com/piotrklosowski96/ebiznes24/commit/27f71e4fe35166da59bc861eafee10251e0eef23)

:white_check_mark: 5.0 Należy dodać konfigurację docker-compose<br>
[Link do commita 5](https://github.com/piotrklosowski96/ebiznes24/commit/f6a73d980895a7ed98a23d1b793eee4a8a591748)

**Wymaganie 3 i 4 zaimplementowałem w tym samym commicie.**

Kod: [Folder z zadaniem 1](https://github.com/piotrklosowski96/ebiznes24/tree/master/Zadanie%201%20-%20Docker)

Link do obrazu na hubie: [Link do hub.docker.com](https://hub.docker.com/layers/piotrklosowski/ebiznes24/latest/images/sha256-8bb50a142475e4689a0c86ad12058ca7a402c434abe0680c84d5aa485d871e13?context=repo)

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/master/Zadanie%201%20-%20Docker/demo)

---

**Zadanie 2 - Scala**

Należy stworzyć aplikację na frameworku Play w Scali 2.

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów<br>
:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy<br>
:white_check_mark: 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD<br>
:white_check_mark: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok<br>
:white_check_mark: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all, show by id (get), update (put), delete (delete), add (post).<br>

[Link do commita z **całą** implementacją](https://github.com/piotrklosowski96/ebiznes24/commit/e8aea0089a062fcf5bb83e045af60f784c23ab40)

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/master/Zadanie%202%20-%20Scala/demo)

---

**Zadanie 3 - Kotlin**

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord<br>
[Link do commita 1](https://github.com/piotrklosowski96/ebiznes24/tree/ae3e95cd7bff8bf824224ada347b649455e25048)

:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota)<br>
[Link do commita 2](https://github.com/piotrklosowski96/ebiznes24/tree/22f648974e2257afdb7f7fab183fd63a0d95eec7)

:white_check_mark: 4.0 Zwróci listę kategorii na określone żądanie użytkownika<br>
[Link do commita 3](https://github.com/piotrklosowski96/ebiznes24/tree/98834633f8ba694cbcacb5adcbf157dcc5b334b1)

:white_check_mark: 4.5 Zwróci listę produktów wg żądanej kategorii<br>
[Link do commita 4](https://github.com/piotrklosowski96/ebiznes24/tree/98834633f8ba694cbcacb5adcbf157dcc5b334b1)

:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenfer, Webex, Skype, Discrod

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/57d010221659d4e1765f5dd414734715ca02e99d/Zadanie%203%20-%20Kotlin/demo)

---

**Zadanie 4 - Go**

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD<br>
[Link do commita 1](https://github.com/piotrklosowski96/ebiznes24/tree/130b2e018c8ea73e18f99667d79701aa70778eaa)<br>
[Link do commita 2](https://github.com/piotrklosowski96/ebiznes24/tree/fae3bce2e73016079773a668544b10975db1b6b8)<br>
[Link do commita 3](https://github.com/piotrklosowski96/ebiznes24/tree/e987095ac632c630a012e9987bb1f6eee2699537)<br>

:white_check_mark: 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)<br>
[Link do commita 4](https://github.com/piotrklosowski96/ebiznes24/tree/e1b12ad438a2b19d8c405d1753c403720f054f1c)

:white_check_mark: 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint<br>
[Link do commita 5](https://github.com/piotrklosowski96/ebiznes24/tree/2b7e6bbc5857d983986f504b7e69573bb4d638b2)

:white_check_mark: 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem<br>
[Link do commita 6](https://github.com/piotrklosowski96/ebiznes24/tree/e7f2e8b48c640c105dd41a37357d33d28bb42cfd)

:white_check_mark: 5.0 pogrupować zapytania w gorm’owe scope'y<br>
[Link do commita 7](https://github.com/piotrklosowski96/ebiznes24/tree/185d7f0979d6a6d41452881e05f54d412aa528eb)

Poprawki zaimplementowałem z małym (~1h) poślizgiem - demo uwzględnia je:

[Link do poprawki 1](https://github.com/piotrklosowski96/ebiznes24/tree/0bac0f2d200f28ebe8d646751638c7a99746f973)<br>
[Link do poprawki 2](https://github.com/piotrklosowski96/ebiznes24/tree/92bdb5c28eb894ad9118d7b7de6cc654097ceb66)<br>
[Link do poprawki 3](https://github.com/piotrklosowski96/ebiznes24/tree/c9e272e2845f52cb24d74fd24ba46149a1e8f0ba)<br>

Demo: [Klik](https://github.com/piotrklosowski96/ebiznes24/tree/57d010221659d4e1765f5dd414734715ca02e99d/Zadanie%204%20-%20Go/demo)

---

# Wszystkie rzeczy związane z backendem od tego momentu były realizowane w ramach jednego projektu i są na Gitlabie.

Do zrealizowania miałem jeszcze projekt w ramach jednego kursy i czasowo nie było dla mnie innej opcji.

Linki:<br>
* [Frontend](https://github.com/piotrklosowski96/pai)
* [Backend](https://gitlab.com/piotr.klosowski96/wpz-2024v2/-/tree/822d1189cd156583b7be43735ceb651e6b5d3539)

**Zadanie 6 - Testy**

Należy stworzyć 20 przypadków testowych w jednym z rozwiązań:

Cypress JS (JS)
Selenium (Kotlin, Python, Java, JS, Go, Scala)
Testy mają w sumie zawierać minimum 50 asercji (3.5). Mają również uruchamiać się na platformie Browserstack (5.0). Proszę pamiętać o stworzeniu darmowego konta via https://education.github.com/pack.

3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium (Kotlin, Python, Java, JS, Go, Scala)

3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50 asercji

:white_check_mark: 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego projektu z minimum 50 asercjami

:white_check_mark: 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z minimum jednym scenariuszem negatywnym per endpoint

5.0 Należy uruchomić testy funkcjonalne na Browserstacku

Demo: [Klik](https://gitlab.com/piotr.klosowski96/wpz-2024v2/-/tree/main/demos/zadanie6?ref_type=heads)

---

**Zadanie 7 - Sonar**

Należy dodać projekt aplikacji klienckiej oraz serwerowej (jeden branch, dwa repozytoria) do Sonara w wersji chmurowej (https://sonarcloud.io/). Należy poprawić aplikacje uzyskując 0 bugów, 0 zapaszków, 0 podatności, 0 błędów bezpieczeństwa. Dodatkowo należy dodać widżety sonarowe do README w repozytorium dane projektu z wynikami.

:white_check_mark: 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w hookach gita

:white_check_mark: 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod aplikacji serwerowej)

:white_check_mark: 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod aplikacji serwerowej)

:white_check_mark: 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa w kodzie w Sonarze (kod aplikacji serwerowej)

5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie aplikacji klienckiej

* https://golangci-lint.run/
* https://github.com/pinterest/ktlint
* https://scalameta.org/scalafmt/docs/installation.html

Demo: [Klik](https://gitlab.com/piotr.klosowski96/wpz-2024v2/-/tree/main/demos/zadanie7?ref_type=heads)

---

**Zadanie 8 - Oauth2**

Należy skonfigurować klienta Oauth2 (4.0). Dane o użytkowniku wraz z tokenem powinny być przechowywane po stronie bazy serwera, a nowy token (inny niż ten od dostawcy) powinien zostać wysłany do klienta (React). Można zastosować mechanizm sesji lub inny dowolny (5.0). Zabronione jest tworzenie klientów bezpośrednio po stronie React'a wyłączając z komunikacji aplikację serwerową, np. wykorzystując auth0.

Prawidłowa komunikacja: react-sewer-dostawca-serwer(via return uri)-react.

:white_check_mark: 3.0 logowanie przez aplikację serwerową (bez Oauth2)

:white_check_mark: 3.5 rejestracja przez aplikację serwerową (bez Oauth2)

:white_check_mark: 4.0 logowanie via Google OAuth2

:white_check_mark: 4.5 logowanie via Facebook lub Github OAuth2

:white_check_mark: 5.0 zapisywanie danych logowania OAuth2 po stronie serwera

Demo: [Klik](https://gitlab.com/piotr.klosowski96/wpz-2024v2/-/tree/main/demos/zadanie8?ref_type=heads)

---

**Zadanie 9 - ChatGPT bot**

[https://gitlab.com/piotr.klosowski96/ollama3.1/-/tree/72ea4beba6c0fca4f60b12b6fa4b31001e981b2d](https://gitlab.com/piotr.klosowski96/ollama3.1/-/tree/72ea4beba6c0fca4f60b12b6fa4b31001e981b2d)

Należy rozszerzyć funkcjonalność wcześniej stworzonego bota. Do niego należy stworzyć aplikajcę frontendową, która połączy się z osobnym serwisem, który przeanalizuje tekst od użytkownika i prześle zapytanie do GPT, a następnie prześle odpowiedź do użytkownika. Cały projekt należy stworzyć w Pythonie.

Dla studentów, którzy nie chcą lub nie mogą korzystać z GPT, zamiast GPT należy wykorzystać LLAMA2 za pomocą narzędzi do wykorzystania LLM lokalnie: https://ollama.com/download/windows

:white_check_mark: 3.0 należy stworzyć po stronie serwerowej osobny serwis do łącznia z chatGPT do usługi chat

3.5 należy stworzyć interfejs frontowy dla użytkownika, który komunikuje się z serwisem; odpowiedzi powinny być wysyałen do frontendowego interfejsu

4.0 stworzyć listę 5 różnych otwarć oraz zamknięć rozmowy

4.5 filtrowanie po zagadnieniach związanych ze sklepem (np. ograniczenie się jedynie do ubrań oraz samego sklepu) do GPT

5.0 filtrowanie odpowiedzi po sentymencie

Demo: [Klik](https://gitlab.com/piotr.klosowski96/wpz-2024v2/-/tree/main/demos/zadanie9?ref_type=heads)

---

**Zadanie 10 - Chmura/CI**

Należy wykorzystać GitHub Actions (dopuszczalne są inne rozwiązania CI) oraz chmurę Azure (dopuszczalne inne chmury), aby zbudować oraz zdeployować aplikację kliencką (frontend) oraz serwerową (backend) jako osobne dwie aplikacje. Należy do tego wykorzystać obrazy dockerowe, a aplikacje powinny działać na kontenerach. Dopuszczalne jest zbudowanie wcześniej aplikacji (jar package) oraz budowanie aplikacji via Github Actions. Należy zwrócić uwagę na zasoby dostępne na chmurze.

:white_check_mark: 3.0 Należy stworzyć odpowiednie instancje po stronie chmury na dockerze

:white_check_mark: 3.5 Stworzyć odpowiedni pipeline w Github Actions do budowania aplikacji (np. via fatjar)

:white_check_mark: 4.0 Dodać notyfikację mailową o zbudowaniu aplikacji

:white_check_mark: 4.5 Dodać krok z deploymentem aplikacji serwerowej oraz klienckiej na chmurę

:white_check_mark: 5.0 Dodać uruchomienie regresyjnych testów automatycznych (funkcjonalnych) jako krok w Actions

Demo: [Klik](https://gitlab.com/piotr.klosowski96/wpz-2024v2/-/tree/main/demos/zadanie10?ref_type=heads)
