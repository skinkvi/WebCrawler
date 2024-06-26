<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Web Crawler Search</title>
    <link rel="stylesheet" href="/static/css/style.css">
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500&display=swap" rel="stylesheet">

</head>
<body>
    <h1>Web Crawler Search</h1>
    <form action="/search" method="post">
        <label for="query">Search query:</label><br>
        <input type="text" id="query" name="query" required autofocus placeholder="Enter search query..."><br>
        <input type="submit" value="Submit">
    </form>
    <button id="clear-results">Clear results</button>

    {{ if .results }}
        <h1>Search results</h1>
        <ul>
            {{ range .results }}
                <li>
                    <a href="{{ .URL }}">{{ .Title }}</a>
                </li>
            {{ end }}
        </ul>
    {{ end }}
    <script>
        document.getElementById('clear-results').addEventListener('click', function() {
            document.getElementById('results-list').innerHTML = '';
        });
    </script>
</body>
</html>
