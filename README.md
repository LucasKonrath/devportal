# My Blog

A simple, elegant blog website built with **Jekyll** and **Markdown**, designed to be hosted on GitHub Pages.

## 🌟 Features

- **Write in Markdown** - No HTML required!
- **Jekyll powered** - Static site generation with built-in support
- **GitHub Pages ready** - Automatic building and deployment
- Clean and modern design
- Fully responsive (works on all devices)
- Easy to customize
- Fast loading times
- SEO optimized

## 🚀 Deploying to GitHub Pages

### Step 1: Create a GitHub Repository

1. Go to [GitHub](https://github.com) and sign in
2. Click the "+" icon in the top right and select "New repository"
3. Name your repository:
   - For a user/org site: `username.github.io` (replace `username` with your GitHub username)
   - For a project site: any name like `my-blog`
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
git commit -m "Initial commit: Jekyll blog setup"

# Add your GitHub repository as remote
git remote add origin https://github.com/YOUR-USERNAME/YOUR-REPO-NAME.git

# Push to GitHub
git branch -M main
git push -u origin main
```

### Step 3: Enable GitHub Pages

1. Go to your repository on GitHub
2. Click on **Settings**
3. Scroll down to **Pages** in the left sidebar
4. Under "Source", select **main** branch
5. Click **Save**
6. Wait a few minutes, then your site will be live at:
   - User/org site: `https://YOUR-USERNAME.github.io/`
   - Project site: `https://YOUR-USERNAME.github.io/YOUR-REPO-NAME/`

### Step 4: Update Configuration (for project sites)

If using a project site (not `username.github.io`), edit `_config.yml`:

```yaml
url: "https://YOUR-USERNAME.github.io"
baseurl: "/YOUR-REPO-NAME"
```

Commit and push the changes.

## 📝 Adding New Blog Posts

Creating new blog posts is easy with Jekyll! Just follow these steps:

### Quick Start

1. Create a new file in the `_posts` directory
2. Name it with the format: `YYYY-MM-DD-title-of-post.md`
3. Add front matter and write your content in Markdown

### Example Post

Create a file: `_posts/2025-11-15-my-new-post.md`

```markdown
---
layout: post
title: "My New Post Title"
date: 2025-11-15
reading_time: 4
excerpt: "A brief description of what this post is about. This will appear on the home page."
---

# This is a heading

Write your content here in **Markdown**. It's simple and powerful!

## Another heading

- Bullet point 1
- Bullet point 2

You can add [links](https://example.com), images, code blocks, and more!

\`\`\`python
def hello():
    print("Hello, World!")
\`\`\`
```

### Front Matter Fields

- `layout`: Always use `post`
- `title`: Your post title (required)
- `date`: Publication date in YYYY-MM-DD format (required)
- `reading_time`: Estimated reading time in minutes (optional)
- `excerpt`: Short description for the home page (optional)

### Markdown Tips

- `# Heading 1` → Main heading
- `## Heading 2` → Subheading
- `**bold text**` → **bold text**
- `*italic text*` → *italic text*
- `[link text](url)` → Hyperlink
- `![alt text](image-url)` → Image
- `` `code` `` → Inline code
- Use triple backticks for code blocks

### Publishing

Just commit and push your new post:

```bash
git add _posts/2025-11-15-my-new-post.md
git commit -m "Add new post"
git push
```

GitHub Pages will automatically rebuild your site!

## 🎨 Customization

### Change Site Settings

Edit `_config.yml` to customize your blog:

```yaml
title: My Awesome Blog          # Your blog title
description: My blog description # Site description
author: Your Name               # Your name
email: your.email@example.com   # Your email
```

### Change Colors and Styles

Edit `assets/css/styles.css` and modify the CSS variables:

```css
:root {
    --primary-color: #2c3e50;    /* Dark blue-gray */
    --secondary-color: #3498db;   /* Blue */
    --accent-color: #e74c3c;      /* Red */
    /* ... other colors ... */
}
```

### Modify Layouts

- Edit `_layouts/default.html` for the overall page structure
- Edit `_layouts/post.html` for individual post layout
- Customize the header and footer sections

## 📁 Project Structure

```
devportal/
├── _config.yml              # Jekyll configuration
├── _layouts/                # Page layouts
│   ├── default.html        # Main layout wrapper
│   └── post.html           # Blog post layout
├── _posts/                  # Your blog posts (Markdown)
│   ├── 2025-11-12-welcome-to-my-blog.md
│   └── 2025-11-10-getting-started-github-pages-jekyll.md
├── assets/
│   └── css/
│       └── styles.css      # All CSS styles
├── index.html              # Home page
├── Gemfile                 # Ruby dependencies
├── .gitignore             # Git ignore rules
└── README.md              # This file
```

## 🔧 Local Development

To test your blog locally before deploying:

### Prerequisites

Install Ruby and Bundler:
- **macOS**: Ruby comes pre-installed. Install bundler with `gem install bundler`
- **Windows**: Install from [RubyInstaller](https://rubyinstaller.org/)
- **Linux**: Use your package manager (e.g., `sudo apt-get install ruby-full`)

### Run Locally

```bash
# Install dependencies (first time only)
bundle install

# Start the Jekyll server
bundle exec jekyll serve

# Or with live reload
bundle exec jekyll serve --livereload
```

Then open `http://localhost:4000` in your browser.

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
