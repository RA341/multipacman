export function showError(message: string): void {
    const overlay = document.getElementById('errorPopupOverlay') as HTMLElement | null;
    const errorMessageElement = document.getElementById('errorMessage') as HTMLParagraphElement | null;
    let closeButton = document.getElementById('closePopupButton') as HTMLButtonElement | null;

    if (overlay && errorMessageElement && closeButton) {
        errorMessageElement.textContent = message;
        overlay.style.display = 'block';

        // Function to close the popup and navigate
        const closePopupAndNavigate = () => {
            if (overlay) {
                overlay.style.display = 'none';
            }
            // Remove event listeners to prevent memory leaks and multiple navigations
            closeButton!.removeEventListener('click', closePopupAndNavigate);
            overlay.removeEventListener('click', closePopupOutsideAndNavigate);

            // Navigate to "/"
            window.location.href = "/";
        };

        // Function to close popup if overlay (outside the content) is clicked
        const closePopupOutsideAndNavigate = (event: MouseEvent) => {
            if (event.target === overlay) { // Check if the click was directly on the overlay
                closePopupAndNavigate();
            }
        };

        // Remove any pre-existing listeners to be safe before adding new ones
        // This is important if showError can be called multiple times while popup might be open
        // (though ideally it shouldn't)
        let oldButton = closeButton.cloneNode(true) as HTMLButtonElement;
        closeButton.parentNode?.replaceChild(oldButton, closeButton);
        closeButton = oldButton; // Re-assign to the new button element

        // let oldOverlay = overlay.cloneNode(false) as HTMLElement; // shallow clone for overlay

        closeButton.addEventListener('click', closePopupAndNavigate);
        overlay.addEventListener('click', closePopupOutsideAndNavigate);

    } else {
        console.error("Popup elements not found in the DOM. Could not display error.");
        alert(message + "\n\nRedirecting to home page.");
        window.location.href = "/";
    }
}

export const delay = (ms: number) => new Promise(res => setTimeout(res, ms));

