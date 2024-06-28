## Challenge:

Create a web scraper that fetches and processes HTML content from a given URL. The scraper should clean up the HTML, fix certain elements, and extract a limited number of specific links from the page. The processed HTML should be saved to a file, and the scraper should be able to traverse pages up to a specified depth.

## Requirements:

### Fetch HTML Content:
- Implement a function that takes a URL as input and fetches the HTML content from that URL.
- Log the size of the fetched HTML content in kilobytes.

### Process HTML Content:
- Use a library  to parse and manipulate the HTML content.

### Clean up the HTML by removing certain elements such as:
- script tags
- Elements with the class .vector-header
- nav tags
- Elements with the ID #p-lang-btn
- Elements with the class .infobox
- Ensure all stylesheet links (link rel="stylesheet") are filtered out

### Save Processed HTML:
- Save the cleaned HTML content to a file in an output directory.
- Ensure the output directory exists; create it if it does not.

### Extract Links:
- Limit the number of extracted links per page to a configurable number (e.g., 10).

### Traversal Depth:
- Implement functionality to traverse links up to a specified depth from the initial page.
- Depth should be a configurable parameter.

### Command-Line Interface:

- The script should be runnable from the command line and take a URL and a depth as an argument.
- If no URL is provided, the script should log an error message and exit.

### Logging:
- Log meaningful messages at various steps, such as when a page is being loaded, the size of the fetched HTML, and the links being processed.

### Example Usage:
- bash
`node script.js https://example.com --depth 2 --per-page 10`

### Additional Considerations:
- Handle edge cases and potential errors gracefully, such as network issues or invalid URLs.




