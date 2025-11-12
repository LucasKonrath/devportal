# My Blog

A technical blog built with **mdBook** (Rust), featuring posts about software development, algorithms, and problem-solving.

## 🌟 Features

- **Write in Markdown** - Simple and powerful
- **mdBook powered** - Fast Rust-based static site generator
- **GitHub Pages ready** - Automatic building and deployment
- **Syntax highlighting** - Beautiful code highlighting out of the box
- **Full-text search** - Built-in search functionality
- **Responsive design** - Works perfectly on all devices
- **Fast loading** - Optimized static site generation

## 🚀 Quick Start

### Prerequisites

Install Rust and mdBook:

```bash
# Install Rust (if not already installed)
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Install mdBook
cargo install mdbook
```

## 🚀 Deploying to GitHub Pages

### Step 1: Create a GitHub Repository

1. Go to [GitHub](https://github.com) and sign in
2. Click the "+" icon in the top right and select "New repository"
3. Name your repository: `devportal` (or any name you prefer)
4. Set it to **Public**
5. Click "Create repository"

### Step 2: Push Your Code to GitHub

Open your terminal and run these commands from your project directory:

```bash
# Initialize git repository (if not already done)
git init

# Add all files
git add .

# Commit your changes
git commit -m "Initial commit: mdBook blog setup"

# Add your GitHub repository as remote
git remote add origin https://github.com/LucasKonrath/devportal.git

# Push to GitHub
git branch -M main
git push -u origin main
```

### Step 3: Enable GitHub Pages

1. Go to your repository on GitHub
2. Click on **Settings**
3. Scroll down to **Pages** in the left sidebar
4. Under "Build and deployment":
   - Source: Select **GitHub Actions**
5. The workflow will automatically run and deploy your site
6. Wait a few minutes, then your site will be live at:
   - `https://LucasKonrath.github.io/devportal/`

That's it! The GitHub Action will automatically build and deploy your mdBook site whenever you push to main.

## 📝 Adding New Blog Posts

Creating new blog posts with mdBook is simple!

### Quick Start

1. Create a new Markdown file in the `src/posts` directory
2. Write your content in Markdown
3. Add an entry to `src/SUMMARY.md`

### Example: Creating a New Post

**Step 1:** Create the file `src/posts/my-new-post.md`

```markdown
# My New Post Title

**November 15, 2025** • 4 min read

Write your content here in **Markdown**. It's simple and powerful!

## A Section

Some content with `code` and more.

\`\`\`rust
fn main() {
    println!("Hello, World!");
}
\`\`\`
```

**Step 2:** Add it to `src/SUMMARY.md`

```markdown
# Blog Posts

- [Welcome to My Blog](./posts/welcome.md)
- [My New Post](./posts/my-new-post.md)  ← Add this line
- [Two Pointers Pattern](./posts/two-pointers.md)
```

**Step 3:** Build and preview

```bash
mdbook build
mdbook serve  # Preview at http://localhost:3000
```

**Step 4:** Publish

```bash
git add src/posts/my-new-post.md src/SUMMARY.md
git commit -m "Add new post"
git push
```

GitHub Actions will automatically rebuild and deploy your site!

## 🎨 Customization

### Change Site Settings

Edit `book.toml` to customize your blog:

```toml
[book]
title = "My Awesome Blog"
authors = ["Your Name"]
description = "Your blog description"
```

### Change Colors and Styles

Edit `theme/custom.css` to modify styles:

```css
:root {
    --blog-primary: #2c3e50;
    --blog-secondary: #3498db;
    --blog-accent: #e74c3c;
}
```

### Modify Structure

- Edit `src/SUMMARY.md` for the table of contents
- Add new sections and organize posts as you like
- Create subdirectories in `src/` for better organization

## 📁 Project Structure

```
devportal/
├── book.toml              # mdBook configuration
├── src/                   # Source content
│   ├── SUMMARY.md        # Table of contents
│   ├── README.md         # Home page
│   ├── about.md          # About page
│   ├── posts/            # Blog posts
│   │   ├── welcome.md
│   │   ├── two-pointers.md
│   │   └── getting-started-mdbook.md
│   └── images/           # Images
├── theme/                # Custom themes
│   └── custom.css       # Custom CSS
├── .github/
│   └── workflows/
│       └── deploy.yml   # GitHub Actions workflow
└── README.md            # This file
```

## 🔧 Local Development

To test your blog locally:

```bash
# Build the book
mdbook build

# Serve locally with auto-reload
mdbook serve

# Open http://localhost:3000 in your browser
```

The site will automatically rebuild when you make changes!

## � Tips

- **Post filenames**: Always use `YYYY-MM-DD-title.md` format in `_posts/` directory
- **Images**: Store in `assets/images/` and reference with `![alt]({{ '/assets/images/photo.jpg' | relative_url }})`
- **Drafts**: Create `_drafts/` folder for unpublished posts (won't appear on site)
- **Custom domains**: Add a `CNAME` file with your domain and configure DNS
- **HTTPS**: Automatically enabled by GitHub Pages
- **Test locally**: Always test with `bundle exec jekyll serve` before pushing
- **Optimization**: Keep image sizes reasonable (use tools like TinyPNG)
- **SEO**: The site uses jekyll-seo-tag for automatic optimization

## 📚 Resources

- [Jekyll Documentation](https://jekyllrb.com/docs/)
- [GitHub Pages Documentation](https://docs.github.com/en/pages)
- [Markdown Guide](https://www.markdownguide.org/)
- [Liquid Template Language](https://shopify.github.io/liquid/)
- [Jekyll Themes](https://jekyllrb.com/docs/themes/)

## 🆘 Troubleshooting

**Site not updating?**
- Check GitHub Actions tab for build errors
- Ensure `_config.yml` has correct `baseurl` for project sites
- Clear browser cache

**Local build failing?**
- Run `bundle update` to update dependencies
- Check Ruby version (Jekyll 4.3 requires Ruby 2.5+)
- Delete `_site/` folder and rebuild

**Posts not showing?**
- Verify filename format: `YYYY-MM-DD-title.md`
- Check that date in filename and front matter match
- Ensure file is in `_posts/` directory

## 🤝 Contributing

Feel free to fork this project and customize it for your own use!

## 📄 License

This project is open source and available under the MIT License.

---

**Happy blogging! 📝✨**
