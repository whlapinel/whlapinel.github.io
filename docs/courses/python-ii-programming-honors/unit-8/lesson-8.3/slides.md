---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

# Lesson 3: HTML & CSS Deep Dive

## Objective

Strengthen understanding of HTML/CSS for creating aesthetically pleasing static sites.

# Semantic HTML Elements

## What are Semantic Elements?

- Elements that convey meaning about their content.

### Examples

- `<header>`: Represents introductory content.
- `<footer>`: Represents footer content.
- `<section>`: Groups related content.
- `<article>`: Represents self-contained content.

### Why Use Semantic Elements?

- Improves accessibility.
- Helps with SEO.
- Makes code more readable.

# CSS Layout Techniques

## Flexbox

- Align and distribute items within a container.
- **Properties to Learn:**
  - `display: flex;`
  - `justify-content`
  - `align-items`

## Grid

- Create two-dimensional layouts.
- **Properties to Learn:**
  - `display: grid;`
  - `grid-template-rows`
  - `grid-template-columns`

# Responsive Design

## Why Responsive Design?

- Adapts layouts to different screen sizes.

# Tools

- **Media Queries:**

  ```css
  @media (max-width: 600px) {
    body {
      font-size: 14px;
    }
  }
  ```

- **Units:**
  - Relative units (e.g., `%`, `em`, `rem`).
  - Flexible layout with `minmax()` in Grid.

# Activity

## Build a Responsive Portfolio Page

### Instructions

1. **HTML Structure**:
   - Use semantic elements like `<header>`, `<section>`, and `<footer>`.
   - Include sections for "About Me," "Projects," and "Contact."

2. **CSS Styling**:
   - Use Flexbox or Grid to create a visually appealing layout.
   - Apply at least one media query to make it responsive.

3. **Enhancements**:
   - Add a navigation bar and links.
   - Use colors, fonts, and spacing to improve aesthetics.

4. **Presentation**:
   - Students will showcase their portfolio page to the class.
