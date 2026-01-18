# Complete Design Token System

## Table of Contents
- [Color System](#color-system)
- [Typography System](#typography-system)
- [Spacing System](#spacing-system)
- [Shadow System](#shadow-system)
- [Animation System](#animation-system)
- [Component Templates](#component-templates)

---

## Color System

### Primary Palette
```css
:root {
  /* Primary - Blue */
  --color-primary-50: #eff6ff;
  --color-primary-100: #dbeafe;
  --color-primary-200: #bfdbfe;
  --color-primary-300: #93c5fd;
  --color-primary-400: #60a5fa;
  --color-primary-500: #3b82f6;  /* Main */
  --color-primary-600: #2563eb;
  --color-primary-700: #1d4ed8;
  --color-primary-800: #1e40af;
  --color-primary-900: #1e3a8a;
}
```

### Neutral Palette
```css
:root {
  --color-gray-50: #f9fafb;
  --color-gray-100: #f3f4f6;
  --color-gray-200: #e5e7eb;
  --color-gray-300: #d1d5db;
  --color-gray-400: #9ca3af;
  --color-gray-500: #6b7280;
  --color-gray-600: #4b5563;
  --color-gray-700: #374151;
  --color-gray-800: #1f2937;
  --color-gray-900: #111827;
}
```

### Semantic Colors
```css
:root {
  /* Success */
  --color-success-light: #dcfce7;
  --color-success-main: #22c55e;
  --color-success-dark: #16a34a;

  /* Warning */
  --color-warning-light: #fef3c7;
  --color-warning-main: #f59e0b;
  --color-warning-dark: #d97706;

  /* Error */
  --color-error-light: #fee2e2;
  --color-error-main: #ef4444;
  --color-error-dark: #dc2626;

  /* Info */
  --color-info-light: #dbeafe;
  --color-info-main: #3b82f6;
  --color-info-dark: #2563eb;
}
```

### Tailwind Mapping
```javascript
// tailwind.config.js
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#eff6ff',
          500: '#3b82f6',
          900: '#1e3a8a',
        },
        success: '#22c55e',
        warning: '#f59e0b',
        error: '#ef4444',
      },
    },
  },
};
```

---

## Typography System

### Font Stack
```css
:root {
  --font-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  --font-mono: 'JetBrains Mono', 'Fira Code', monospace;
}
```

### Type Scale
| Token | Size | Weight | Line Height | Use Case |
|-------|------|--------|-------------|----------|
| `display` | 3rem (48px) | 700 | 1.1 | Hero headlines |
| `h1` | 2.25rem (36px) | 700 | 1.2 | Page titles |
| `h2` | 1.875rem (30px) | 600 | 1.3 | Section headers |
| `h3` | 1.5rem (24px) | 600 | 1.4 | Subsections |
| `h4` | 1.25rem (20px) | 500 | 1.4 | Card titles |
| `body-lg` | 1.125rem (18px) | 400 | 1.6 | Lead paragraphs |
| `body` | 1rem (16px) | 400 | 1.5 | Body text |
| `body-sm` | 0.875rem (14px) | 400 | 1.5 | Secondary text |
| `caption` | 0.75rem (12px) | 400 | 1.4 | Labels, hints |

### CSS Implementation
```css
.text-display { font-size: 3rem; font-weight: 700; line-height: 1.1; }
.text-h1 { font-size: 2.25rem; font-weight: 700; line-height: 1.2; }
.text-h2 { font-size: 1.875rem; font-weight: 600; line-height: 1.3; }
.text-h3 { font-size: 1.5rem; font-weight: 600; line-height: 1.4; }
.text-body { font-size: 1rem; font-weight: 400; line-height: 1.5; }
.text-caption { font-size: 0.75rem; font-weight: 400; line-height: 1.4; }
```

---

## Spacing System

### Base Unit: 4px

| Token | Value | Tailwind | Use Case |
|-------|-------|----------|----------|
| `space-0` | 0 | `p-0` | Reset |
| `space-1` | 4px | `p-1` | Tight inline |
| `space-2` | 8px | `p-2` | Icon gaps |
| `space-3` | 12px | `p-3` | Button padding |
| `space-4` | 16px | `p-4` | Card padding |
| `space-5` | 20px | `p-5` | Section gap |
| `space-6` | 24px | `p-6` | Content blocks |
| `space-8` | 32px | `p-8` | Large sections |
| `space-10` | 40px | `p-10` | Hero spacing |
| `space-12` | 48px | `p-12` | Page margins |
| `space-16` | 64px | `p-16` | Major sections |

### Component Spacing Guidelines
```
Card:        padding: space-4 (16px)
Modal:       padding: space-6 (24px)
Form Field:  gap: space-2 (8px)
Button:      padding: space-3 space-4 (12px 16px)
Input:       padding: space-2 space-3 (8px 12px)
```

---

## Shadow System

```css
:root {
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
  --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
  --shadow-lg: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1);
  --shadow-xl: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);
}
```

| Token | Use Case |
|-------|----------|
| `shadow-sm` | Subtle depth (cards at rest) |
| `shadow-md` | Elevated cards, dropdowns |
| `shadow-lg` | Modals, popovers |
| `shadow-xl` | Floating action buttons |

---

## Animation System

### Duration
```css
:root {
  --duration-fast: 150ms;
  --duration-normal: 200ms;
  --duration-slow: 300ms;
}
```

### Easing
```css
:root {
  --ease-in: cubic-bezier(0.4, 0, 1, 1);
  --ease-out: cubic-bezier(0, 0, 0.2, 1);
  --ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);
}
```

### Common Transitions
```css
.transition-colors { transition: color, background-color var(--duration-fast) var(--ease-in-out); }
.transition-opacity { transition: opacity var(--duration-normal) var(--ease-in-out); }
.transition-transform { transition: transform var(--duration-fast) var(--ease-out); }
.transition-all { transition: all var(--duration-normal) var(--ease-in-out); }
```

---

## Component Templates

### Button
```markdown
## Button

| Property | Default | Hover | Active | Disabled |
|----------|---------|-------|--------|----------|
| Background | primary-500 | primary-600 | primary-700 | gray-300 |
| Text | white | white | white | gray-500 |
| Scale | 1 | 1 | 0.98 | 1 |
| Opacity | 1 | 1 | 1 | 0.5 |
| Cursor | pointer | pointer | pointer | not-allowed |

Sizes:
- sm: h-8, px-3, text-sm
- md: h-10, px-4, text-base
- lg: h-12, px-6, text-lg
```

### Input
```markdown
## Input

| Property | Default | Focus | Error | Disabled |
|----------|---------|-------|-------|----------|
| Border | gray-300 | primary-500 | error | gray-200 |
| Background | white | white | error-light | gray-100 |
| Ring | none | 2px primary | 2px error | none |

Structure:
- Label: text-sm, text-gray-700, mb-1
- Input: h-10, px-3, rounded-md
- Helper: text-xs, text-gray-500, mt-1
- Error: text-xs, text-error, mt-1
```

### Card
```markdown
## Card

| Variant | Background | Border | Shadow |
|---------|------------|--------|--------|
| Default | white | gray-200 | shadow-sm |
| Elevated | white | none | shadow-md |
| Outlined | transparent | gray-300 | none |

Structure:
- Container: rounded-lg, p-4
- Header: pb-4, border-b (optional)
- Body: py-4
- Footer: pt-4, border-t (optional)
```
