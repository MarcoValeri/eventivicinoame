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