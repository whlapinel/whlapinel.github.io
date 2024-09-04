(() => {
    console.log("Hello World");
    document.addEventListener("htmx:afterOnLoad", () => {
        console.log("event triggered")
        const pageTitles = ["Home", "About", "Contact", "Blog"];
        const currLocation = window.location.pathname;
        console.log(currLocation);
        for (const title of pageTitles) {
            if (currLocation.includes(title.toLowerCase())) {
                console.log(title);
                const titleElement = document.querySelector(`#nav-${title.toLowerCase()}`);
                const selectedElement = document.querySelector(`.highlighted`)
                selectedElement?.classList.remove("bg-green-700");
                selectedElement?.classList.remove("highlighted");
                console.log(titleElement);
                titleElement?.classList.toggle("bg-green-700");
                titleElement?.classList.add(".highlighted");
            }
        }
    })
})()