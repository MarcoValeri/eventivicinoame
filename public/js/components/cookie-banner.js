{/* <div id="cookie-banner-component" class="cookie-banner">
    <div class="cookie-banner__container-logo">
        <img loading="lazy" class="cookie-banner__logo" src="/public/images/eventi-vicino-a-me-logo.png" alt="Logo Eventi Vicino A Me, una ape che vola">
        <h3 class="h4">Ci teniamo alla tua privacy</h3>
    </div>
    <div class="cookie-banner__container-content">
        <p class="p">
            Eventi Vicino A Me utilizza i cookies per migliorare la tua esperienza di navigazione e mostrarti contenuti personalizzati. Cliccando su "Accetta", acconsenti all'uso di tutti i cookies. Puoi leggere la nostra informativa sulla privacy selezionado l'opzione "Scopri di più" per maggiori informazioni.
        </p>
    </div>
    <div class="cookie-banner__container-buttons">
        <button id="cookie-banner-button-acept" class="cookie-banner__button button">Accetta</button>
        <a id="cookie-banner-button-discover" class="cookie-banner__button button" href="/page/privacy-policy">Scopri di più</a>
        <button id="cookie-banner-button-reject" class="cookie-banner__button button">Rifiuta</button>
    </div>
</div> */}

const setNewCookie = (cookieName, cookieValue, cookieDaysToExpire) => {
    let expires = "";
    if (cookieDaysToExpire) {
        const date = new Date();
        date.setTime(date.getTime() + (cookieDaysToExpire * 24 * 60 * 60 * 1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = cookieName + "=" + (cookieValue || "") + expires + "; path=/";
}

const checkCookieBanner = (getValue) => {
    const cookies = document.cookie.split(";");

    for (let i = 0; i < cookies.length; i++) {
        let cookie = cookies[i].trim();

        if (cookie.startsWith("cookieBanner=")) {
            let cookieValue = cookie.substring("cookieBanner=".length);
            if (cookieValue === getValue) {
                return true;
            } else {
                return false;
            }
        }
    }

    return false;
}

const cookieBanner = () => {
    const cookieBannerComponent = document.getElementById('cookie-banner-component');
    const cookieBannerButtonAcept = document.getElementById('cookie-banner-button-acept');
    const cookieBannerButtonDiscover = document.getElementById('cookie-banner-button-discover');
    const cookieBannerButtonReject = document.getElementById('cookie-banner-button-reject');

    if (checkCookieBanner("accepted") || checkCookieBanner("rejected")) {
        cookieBannerComponent.style.display = "none";
    } else {
        cookieBannerComponent.style.display = "block";
    }

    // Click event
    if (cookieBannerButtonAcept) {
        cookieBannerButtonAcept.addEventListener('click', () => {
            cookieBannerComponent.style.display = "none";
            setNewCookie("cookieBanner", "accepted", 30);
        })
    }

    // Click event
    if (cookieBannerButtonDiscover) {
        cookieBannerButtonDiscover.addEventListener('click', () => {
            cookieBannerComponent.style.display = "none";
            setNewCookie("cookieBanner", "accepted", 30);
        })
    }

    // Click event
    if (cookieBannerButtonReject) {
        cookieBannerButtonReject.addEventListener('click', () => {
            cookieBannerComponent.style.display = "none";
            setNewCookie("cookieBanner", "rejected", 30);
        })
    }
}

export { cookieBanner };