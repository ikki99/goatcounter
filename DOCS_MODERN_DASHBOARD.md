# Modern Dashboard Technical Documentation (Lite Version)

## 1. Overview
The Modern Dashboard is a React-based SPA (Single Page Application) embedded within GoatCounter. It is designed for high performance and visual excellence, utilizing a customized API layer for real-time statistics.

## 2. Performance Optimizations

### 2.1 Parallel Data Fetching
In the standard dashboard, widgets load sequentially. The Lite version refactors this using `Promise.all` patterns:
- **Tiered Loading**: Requests are grouped into tiers (KPIs, Charts, Lists).
- **Internal Parallelism**: Each tier fires all its API calls simultaneously.
- **KPI Concurrency**: Primary metrics and trend data are fetched in parallel to minimize "Time to Interactive".

### 2.2 Rate Limit Tuning
- **Constant**: `RATE_LIMIT_MS`
- **Standard**: 500ms
- **Lite**: 150ms
The reduction in throttle delay allows the parallel request burst to complete significantly faster while staying within the backend's safety margins.

## 3. Domain Filtering Logic

### 3.1 Subdomain Extraction
The dashboard automatically detects subdomains when cross-domain tracking is used.
- **Backend**: `handlers/api.go` -> `trackedDomains`
- **Algorithm**: 
  1. Query `hit_counts` joined with `paths` to find all active paths for the site.
  2. Parse the first component of each path (e.g., `/sub.example.com/page` -> `sub.example.com`).
  3. Filter out common file extensions to ensure only valid domains are captured.
  4. Exclude the primary site domain (Cname) to prevent redundancy with the "All" filter.

### 3.2 Frontend Integration
- **State Management**: `filter` state is synchronized with the URL search parameters.
- **UI**: A dynamic button group is rendered at the top of the dashboard.
- **Behavior**: Clicking "All" removes the filter. Clicking a domain updates all widgets in parallel to show data only for that domain.

## 4. Deployment
The dashboard is embedded using Go's `embed` package.
- **Source**: `handlers/modern_assets_embedded/modern.html`
- **Sync**: Changes must be copied to `public/modern.html` for development consistency.
- **Build**: Use `build.ps1` to recompile the binary with the updated assets.
