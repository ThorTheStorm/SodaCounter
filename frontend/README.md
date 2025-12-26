# Soda Inventory Frontend

A modern, dark-mode frontend for the Soda Counter REST API.

## Features

- 🌙 Beautiful dark mode interface
- 📊 Real-time inventory statistics
- 🥤 Visual soda icons based on soda type
- ✨ Smooth animations and transitions
- 📱 Fully responsive design
- ⚡ Fast and lightweight

## Setup Instructions

### 1. Configure CORS in Your Go REST API

Before using the frontend, you need to enable CORS in your Go backend. Add this to your `main-rest.go`:

```go
package main

import (
    "net/http"
    // ... your other imports
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

func main() {
    // ... your existing setup code
    
    // Wrap your router with CORS middleware
    http.ListenAndServe(":8080", corsMiddleware(yourRouter))
}
```

### 2. Update API Configuration

Edit `config.js` and update the `BASE_URL` to match your Go REST API:

```javascript
const API_CONFIG = {
    BASE_URL: 'http://localhost:8080',  // Change this to your API URL
    // ...
};
```

### 3. Serve the Frontend

You have several options:

#### Option A: Using Python (Simplest)
```bash
cd frontend
python -m http.server 3000
```
Then open http://localhost:3000

#### Option B: Using Node.js http-server
```bash
npm install -g http-server
cd frontend
http-server -p 3000
```

#### Option C: Using VS Code Live Server Extension
1. Install "Live Server" extension in VS Code
2. Right-click on `index.html`
3. Select "Open with Live Server"

#### Option D: Using Go (if you prefer)
Add this to your Go REST API to serve the frontend:

```go
// Serve the frontend
fs := http.FileServer(http.Dir("../frontend"))
http.Handle("/", fs)
```

### 4. Start Using the App

1. Make sure your Go REST API is running on port 8080
2. Open the frontend in your browser
3. Start managing your soda inventory!

## API Endpoints Expected

The frontend expects these endpoints from your REST API:

- `GET /sodas` - Get all sodas
- `POST /sodas` - Create a new soda
- `GET /sodas/:id` - Get a specific soda
- `PUT /sodas/:id` - Update a soda
- `DELETE /sodas/:id` - Delete a soda

## Customization

### Adding More Soda Icons

Edit `config.js` and add entries to the `SODA_ICONS` object:

```javascript
const SODA_ICONS = {
    'coca-cola': '🥤',
    'your-soda-name': '🎨',  // Add your custom icon
    // ...
};
```

### Changing Colors

Edit `styles.css` and modify the CSS variables in `:root`:

```css
:root {
    --accent-primary: #3b82f6;  /* Change primary color */
    --accent-success: #10b981;  /* Change success color */
    /* ... */
}
```

## Troubleshooting

### "Failed to load sodas" error
- Make sure your Go REST API is running
- Check that the `BASE_URL` in `config.js` is correct
- Verify CORS is properly configured in your backend

### Can't connect to API
- Check if the API is running on the correct port
- Make sure there are no firewall issues
- Try accessing the API directly in your browser

## Browser Support

- Chrome/Edge (latest)
- Firefox (latest)
- Safari (latest)
- Opera (latest)

## Technologies Used

- Vanilla JavaScript (no frameworks!)
- Modern CSS3 with animations
- Fetch API for HTTP requests
- CSS Grid and Flexbox for layout
