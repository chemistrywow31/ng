# 勞工們

## The Orchestrator (原 Tech Lead / EM)
```text
   Agent 類型： Controller / Planner

核心技能：

Task Decomposition: 將 PM 的需求拆解為具體的、無相依性的開發任務序列 (DAG)。

Reviewer: 檢查 Coder Agent 回傳的結果是否符合 JSON 格式，是否通過了 QA Agent 的測試。

Context Manager: 決定傳送多少「上下文」給下一個 Agent（避免 Token 爆炸）。 
```


## The Product Definer (原 PM)
```text
   Agent 類型： Refiner

核心技能：

Disambiguation: 當用戶輸入模糊指令時，反問用戶以釐清需求。

Spec Generation: 產出結構化的 PRD (Markdown 格式)，包含 User Stories 與 Acceptance Criteria。
```

## The Blueprint Keeper (原 架構師)
```text
Agent 類型： Architect / Memory Holder

核心技能：

Tech Stack Constraint: 強制規定使用的語言、框架版本（防止前端用 React 18，後端給了 jQuery 的寫法）。

File Structure Management: 維護專案的檔案目錄結構樹。

Interface Definition: 定義 API 介面 (OpenAPI Spec) 與資料庫 Schema。
```

## The Builders (原 資深前/後端)
```text
Agent 類型： Executor / Coder

核心技能：

Implementation: 根據輸入的 Spec 與架構限制，輸出可執行的程式碼區塊。

Self-Reflection: 在輸出前自我檢查語法錯誤。

註記： 對 Agent 來說，「資深」意味著 System Prompt 裡包含了 "Expert in [Language], strictly follow Clean Code principles, prioritize performance..." 等高權重指令。
```

## The Critics (原 QA / Security)
```text
Agent 類型： Validator

核心技能：

Static Analysis: 模擬 Compiler 或 Linter 檢查代碼。

Edge Case Generation: 針對 Coder 的產出，思考 3 個可能導致崩潰的極端案例並要求修正。

Security Audit: 掃描常見漏洞 (Injection, XSS) 模式。
```

## The Experience Designer (UI/UX Agent)
```text
Agent 類型： Translator / Specifier (轉譯者 / 規格制定者)

在流程中的位置： 介於 PM (Product Definer) 與 Frontend Builder 之間。

核心任務： 將文字需求轉化為「視覺結構」與「互動邏輯」，確保前端工程師不需要「猜測」樣式。

Interaction Flow Design (UX 邏輯設計):

技能描述： 根據 PM 的 User Story，設計使用者的操作路徑。

輸出形式： Mermaid Flowchart 或文字描述的步驟 (Step-by-Step Interaction)。

關鍵細節： 必須定義「Happy Path」以外的狀態（例如：Loading 狀態、錯誤訊息顯示位置、空資料狀態 Empty State）。

Agent 思考範例：「當使用者點擊登入後，按鈕應變為 Loading 狀態 (Disabled)，若 API 回傳 401，則在 Input 下方顯示紅字錯誤訊息。」

Structural Wireframing (結構化線框圖):

技能描述： 不畫圖，而是用文字或偽代碼 (Pseudo-HTML) 定義頁面的 DOM 結構層級。

輸出形式： Component Tree JSON 或 Semantic HTML 結構草稿。

Agent 思考範例：「頁面佈局為：Header (Logo + Nav), Main (Hero Section + Feature Grid), Footer。Hero Section 包含 H1 標題、P 副標題、CTA 按鈕。」

Design System Enforcement (視覺規範與原子化設計):

技能描述： 這是最重要的部分。它負責將抽象的「好看」轉化為具體的 CSS/Tailwind Class 或 Design Tokens。

輸出形式： JSON Design Tokens (Colors, Typography, Spacing) 或直接指定 UI Library (如 MUI, AntD, Shadcn/ui) 的組件名稱。

Agent 思考範例：「主色調為 #3B82F6 (Blue-500)，按鈕圓角使用 rounded-lg，間距統一使用 p-4，字體使用 Inter。請前端使用 Button Component，variant 設定為 outline。」

Accessibility (A11y) Guard (無障礙守門員):

技能描述： 確保設計符合無障礙標準。

輸出形式： 在規格中標註 aria-label、對比度要求、鍵盤導航順序。
```

