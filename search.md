

# datatools

<link href="./pagefind/pagefind-ui.css" rel="stylesheet">
<script src="./pagefind/pagefind-ui.js" type="text/javascript"></script>
<div id="search"></div>
<script>
const u = URL.parse(window.location.href);
const basePath = u.pathname.replace(/search.html$/g, '');

// Function to extract query parameters from the URL
function getQueryParam(name) {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get(name);
}

// Extract the query parameter
const searchQuery = getQueryParam('q');

window.addEventListener('DOMContentLoaded', (event) => {
    const searchUI = new PagefindUI({ 
            element: "#search",
            baseUrl: basePath
    });
    if (searchQuery) {
        searchUI.triggerSearch(searchQuery);
    }
});
</script>
