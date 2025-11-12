# Quick Start Guide - mdBook Blog

## ✅ Your blog is now powered by mdBook!

### What Changed

- **Framework**: Jekyll → mdBook (Rust-based)
- **Content Location**: `_posts/` → `src/posts/`
- **Configuration**: `_config.yml` → `book.toml`
- **Syntax Highlighting**: Built-in with mdBook (no extra setup needed)
- **Build Time**: Much faster! ⚡

### Local Preview

Your blog is currently running at: **http://127.0.0.1:3000**

```bash
# Start local server (already running)
mdbook serve

# Build static files
mdbook build
```

### Project Structure

```
devportal/
├── book.toml           # Configuration
├── src/                # Your content
│   ├── SUMMARY.md     # Table of contents (sidebar)
│   ├── README.md      # Home page
│   ├── about.md       # About page
│   ├── posts/         # Blog posts
│   │   ├── welcome.md
│   │   ├── getting-started-mdbook.md
│   │   └── two-pointers.md
│   └── images/        # Your images
├── theme/
│   └── custom.css     # Custom styling
└── book/              # Generated HTML (git ignored)
```

### Adding a New Post

1. **Create file**: `src/posts/my-post.md`
   ```markdown
   # My Post Title
   
   **November 15, 2025** • 5 min read
   
   Your content here...
   ```

2. **Update SUMMARY.md**: Add to `src/SUMMARY.md`
   ```markdown
   - [My Post Title](./posts/my-post.md)
   ```

3. **Preview**: Changes appear instantly at http://127.0.0.1:3000

4. **Deploy**: `git add . && git commit -m "Add post" && git push`

### Deploy to GitHub Pages

```bash
# Commit everything
git add .
git commit -m "Convert to mdBook"

# Push to GitHub
git push origin main
```

Then:
1. Go to GitHub repo → Settings → Pages
2. Under "Build and deployment", select **GitHub Actions**
3. Your site will be live at: https://LucasKonrath.github.io/devportal/

The GitHub Action (`.github/workflows/deploy.yml`) handles everything automatically!

### Features

✅ **Syntax highlighting** - Automatic for all languages
✅ **Search** - Full-text search built-in
✅ **Mobile responsive** - Works perfectly on all devices
✅ **Fast builds** - Rust-powered performance
✅ **Table of contents** - Auto-generated from SUMMARY.md
✅ **Print/PDF** - Built-in print stylesheet
✅ **Code playgrounds** - Interactive code blocks (configurable)

### Tips

- Use `mdbook serve --open` to auto-open browser
- Use `mdbook watch` to rebuild on changes without serving
- Use `mdbook test` to test code samples in your posts
- Edit `theme/custom.css` to customize colors and styling

### Customization

Edit `book.toml`:
```toml
[book]
title = "Your Blog Name"
authors = ["Your Name"]
description = "Your description"
```

Edit `theme/custom.css`:
```css
:root {
    --blog-primary: #your-color;
    --blog-secondary: #your-color;
    --blog-accent: #your-color;
}
```

### Resources

- [mdBook User Guide](https://rust-lang.github.io/mdBook/)
- [Markdown Guide](https://www.markdownguide.org/)
- [Your local preview](http://127.0.0.1:3000)

Happy blogging with Rust! 🦀📝
