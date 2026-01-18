# Frontend Component Patterns

## Table of Contents
- [React Patterns](#react-patterns)
- [Vue 3 Patterns](#vue-3-patterns)
- [Custom Hooks / Composables](#custom-hooks--composables)
- [Form Patterns](#form-patterns)
- [Error Handling](#error-handling)

---

## React Patterns

### Base Component Structure
```tsx
// components/Card.tsx
import { FC, ReactNode } from 'react';
import { cn } from '@/lib/utils';

interface CardProps {
  children: ReactNode;
  className?: string;
  variant?: 'default' | 'elevated' | 'outlined';
}

export const Card: FC<CardProps> = ({
  children,
  className,
  variant = 'default'
}) => {
  const variants = {
    default: 'bg-white rounded-lg p-4',
    elevated: 'bg-white rounded-lg p-4 shadow-lg',
    outlined: 'bg-white rounded-lg p-4 border border-gray-200',
  };

  return (
    <div className={cn(variants[variant], className)}>
      {children}
    </div>
  );
};
```

### Compound Component Pattern
```tsx
// components/Tabs/index.tsx
import { createContext, useContext, useState, ReactNode, FC } from 'react';

interface TabsContextType {
  activeTab: string;
  setActiveTab: (id: string) => void;
}

const TabsContext = createContext<TabsContextType | null>(null);

const useTabs = () => {
  const context = useContext(TabsContext);
  if (!context) throw new Error('useTabs must be used within Tabs');
  return context;
};

interface TabsProps {
  children: ReactNode;
  defaultTab: string;
}

export const Tabs: FC<TabsProps> & {
  List: FC<{ children: ReactNode }>;
  Tab: FC<{ id: string; children: ReactNode }>;
  Panels: FC<{ children: ReactNode }>;
  Panel: FC<{ id: string; children: ReactNode }>;
} = ({ children, defaultTab }) => {
  const [activeTab, setActiveTab] = useState(defaultTab);
  return (
    <TabsContext.Provider value={{ activeTab, setActiveTab }}>
      {children}
    </TabsContext.Provider>
  );
};

Tabs.List = ({ children }) => (
  <div className="flex gap-2 border-b">{children}</div>
);

Tabs.Tab = ({ id, children }) => {
  const { activeTab, setActiveTab } = useTabs();
  return (
    <button
      onClick={() => setActiveTab(id)}
      className={cn(
        'px-4 py-2 -mb-px',
        activeTab === id && 'border-b-2 border-blue-500'
      )}
    >
      {children}
    </button>
  );
};

Tabs.Panels = ({ children }) => <div className="py-4">{children}</div>;

Tabs.Panel = ({ id, children }) => {
  const { activeTab } = useTabs();
  return activeTab === id ? <>{children}</> : null;
};
```

### Render Props Pattern
```tsx
// components/DataFetcher.tsx
interface DataFetcherProps<T> {
  url: string;
  children: (props: {
    data: T | null;
    loading: boolean;
    error: Error | null;
  }) => ReactNode;
}

export function DataFetcher<T>({ url, children }: DataFetcherProps<T>) {
  const { data, error, isLoading } = useSWR<T>(url, fetcher);

  return (
    <>
      {children({
        data: data ?? null,
        loading: isLoading,
        error: error ?? null,
      })}
    </>
  );
}

// Usage
<DataFetcher<User[]> url="/api/users">
  {({ data, loading, error }) => {
    if (loading) return <Skeleton />;
    if (error) return <ErrorState />;
    return <UserList users={data!} />;
  }}
</DataFetcher>
```

---

## Vue 3 Patterns

### Base Component (Composition API)
```vue
<!-- components/Card.vue -->
<script setup lang="ts">
interface Props {
  variant?: 'default' | 'elevated' | 'outlined';
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'default',
});

const variantClasses = {
  default: 'bg-white rounded-lg p-4',
  elevated: 'bg-white rounded-lg p-4 shadow-lg',
  outlined: 'bg-white rounded-lg p-4 border border-gray-200',
};
</script>

<template>
  <div :class="variantClasses[props.variant]">
    <slot />
  </div>
</template>
```

### Provide/Inject Pattern
```vue
<!-- components/Tabs/Tabs.vue -->
<script setup lang="ts">
import { provide, ref } from 'vue';

const props = defineProps<{ defaultTab: string }>();
const activeTab = ref(props.defaultTab);

provide('tabs', {
  activeTab,
  setActiveTab: (id: string) => { activeTab.value = id; },
});
</script>

<template>
  <div><slot /></div>
</template>

<!-- components/Tabs/Tab.vue -->
<script setup lang="ts">
import { inject } from 'vue';

const props = defineProps<{ id: string }>();
const { activeTab, setActiveTab } = inject('tabs')!;
</script>

<template>
  <button
    @click="setActiveTab(props.id)"
    :class="{ 'border-b-2 border-blue-500': activeTab === props.id }"
  >
    <slot />
  </button>
</template>
```

---

## Custom Hooks / Composables

### React: useAsync Hook
```tsx
// hooks/useAsync.ts
import { useState, useCallback } from 'react';

interface AsyncState<T> {
  data: T | null;
  loading: boolean;
  error: Error | null;
}

export function useAsync<T, Args extends unknown[]>(
  asyncFn: (...args: Args) => Promise<T>
) {
  const [state, setState] = useState<AsyncState<T>>({
    data: null,
    loading: false,
    error: null,
  });

  const execute = useCallback(async (...args: Args) => {
    setState({ data: null, loading: true, error: null });
    try {
      const data = await asyncFn(...args);
      setState({ data, loading: false, error: null });
      return data;
    } catch (error) {
      setState({ data: null, loading: false, error: error as Error });
      throw error;
    }
  }, [asyncFn]);

  return { ...state, execute };
}
```

### Vue: useAsync Composable
```ts
// composables/useAsync.ts
import { ref, Ref } from 'vue';

export function useAsync<T, Args extends unknown[]>(
  asyncFn: (...args: Args) => Promise<T>
) {
  const data: Ref<T | null> = ref(null);
  const loading = ref(false);
  const error: Ref<Error | null> = ref(null);

  const execute = async (...args: Args) => {
    loading.value = true;
    error.value = null;
    try {
      data.value = await asyncFn(...args);
      return data.value;
    } catch (e) {
      error.value = e as Error;
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return { data, loading, error, execute };
}
```

### useLocalStorage Hook
```tsx
// hooks/useLocalStorage.ts
import { useState, useEffect } from 'react';

export function useLocalStorage<T>(key: string, initialValue: T) {
  const [value, setValue] = useState<T>(() => {
    if (typeof window === 'undefined') return initialValue;
    try {
      const item = window.localStorage.getItem(key);
      return item ? JSON.parse(item) : initialValue;
    } catch {
      return initialValue;
    }
  });

  useEffect(() => {
    window.localStorage.setItem(key, JSON.stringify(value));
  }, [key, value]);

  return [value, setValue] as const;
}
```

---

## Form Patterns

### Controlled Form with Validation
```tsx
// components/LoginForm.tsx
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';

const schema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string().min(8, 'Min 8 characters'),
});

type FormData = z.infer<typeof schema>;

export const LoginForm: FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  });

  const onSubmit = async (data: FormData) => {
    // API call
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div>
        <input {...register('email')} placeholder="Email" />
        {errors.email && <span className="text-red-500">{errors.email.message}</span>}
      </div>
      <div>
        <input {...register('password')} type="password" placeholder="Password" />
        {errors.password && <span className="text-red-500">{errors.password.message}</span>}
      </div>
      <button type="submit" disabled={isSubmitting}>
        {isSubmitting ? 'Loading...' : 'Login'}
      </button>
    </form>
  );
};
```

---

## Error Handling

### Error Boundary
```tsx
// components/ErrorBoundary.tsx
import { Component, ErrorInfo, ReactNode } from 'react';

interface Props {
  children: ReactNode;
  fallback: ReactNode;
}

interface State {
  hasError: boolean;
}

export class ErrorBoundary extends Component<Props, State> {
  state: State = { hasError: false };

  static getDerivedStateFromError(): State {
    return { hasError: true };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    console.error('ErrorBoundary caught:', error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback;
    }
    return this.props.children;
  }
}
```

### Async Error Handling
```tsx
// Safe API wrapper
async function safeFetch<T>(url: string): Promise<{ data: T | null; error: string | null }> {
  try {
    const res = await fetch(url);
    if (!res.ok) {
      return { data: null, error: `HTTP ${res.status}` };
    }
    const data = await res.json();
    return { data, error: null };
  } catch (e) {
    return { data: null, error: (e as Error).message };
  }
}
```
