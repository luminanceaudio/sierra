import React, {
  createContext,
  useContext,
  useEffect,
  useMemo,
  useReducer,
  useState,
} from 'react';
import { ItemData } from './data';

export type FocusElement = 'item' | 'editor';

type AppState = {
  focusedItemIndex: { number: number; element: FocusElement };
  items: ItemData[];
};

const initialState: AppState = {
  focusedItemIndex: { number: -1, element: 'editor' },
  items: [],
};

type Action =
  | {
      type: 'addItem';
      item: ItemData;
      index?: number;
    }
  | {
      type: 'deleteItem';
      index: number;
    }
  | {
      type: 'setItem';
      item: ItemData;
      index: number;
    }
  | {
      type: 'loadState';
      state: AppState;
    }
  | {
      type: 'setFocusedItemIndex';
      index?: number;
      element?: FocusElement;
    };

type Dispatch = (action: Action) => void;

const AppContext = createContext<
  { state: AppState; dispatch: Dispatch } | undefined
>(undefined);

function appReducer(state: AppState, action: Action): AppState {
  switch (action.type) {
    case 'addItem': {
      const newItems = [...state.items];
      if (action.index === undefined) {
        newItems.push(action.item);
      } else {
        newItems.splice(action.index, 0, newItems[action.index]);
        newItems[action.index + 1] = action.item;
      }

      return { ...state, items: newItems };
    }

    case 'deleteItem': {
      let newItems = [...state.items];
      newItems = newItems.filter(
        (v: ItemData, index: number) => index !== action.index,
      );
      return { ...state, items: newItems };
    }

    case 'setItem': {
      const newItems = [...state.items];
      newItems[action.index] = action.item;
      return { ...state, items: newItems };
    }

    case 'loadState': {
      return { ...action.state };
    }

    case 'setFocusedItemIndex': {
      return {
        ...state,
        focusedItemIndex: {
          number: action.index ?? state.focusedItemIndex.number,
          element: action.element ?? state.focusedItemIndex.element,
        },
      };
    }

    default: {
      throw new Error(`Unhandled action: ${action}`);
    }
  }
}

type AppProviderProps = {
  children: React.ReactNode;
};

export function AppProvider({ children }: AppProviderProps) {
  const [state, dispatch] = useReducer(appReducer, initialState);
  const [timeoutId, setTimeoutId] = useState<number | undefined>(undefined);

  const value = useMemo(() => {
    return { state, dispatch };
  }, [state, dispatch]);

  // Load state
  useEffect(() => {
    window.electron.ipcRenderer.once(
      'app-provider.load-state',
      (loadedStateJsonStr) => {
        dispatch({
          type: 'loadState',
          state: JSON.parse(loadedStateJsonStr as string),
        });
      },
    );

    console.log('load state from disk');
    window.electron.ipcRenderer.sendMessage('app-provider.load-state');
  }, []);

  // Save state on change, limited to 1 update a second
  useEffect(() => {
    if (timeoutId) {
      clearTimeout(timeoutId);
    }

    const newTimeoutId = setTimeout(() => {
      window.electron.ipcRenderer.sendMessage(
        'app-provider.save-state',
        JSON.stringify(state),
      );
    }, 1000) as unknown as number; // In renderer process, this returns an ID

    setTimeoutId(newTimeoutId);

    // This useEffect cleanup function will be called
    // when the component unmounts or when the state changes
    return () => {
      clearTimeout(newTimeoutId);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [state]);

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
}

function useApp(): [AppState, Dispatch] {
  const context = useContext(AppContext);
  if (context === undefined) {
    throw new Error('useApp must be used within a AuthContext');
  }
  return [context.state, context.dispatch];
}

export function useItems(): {
  items: ItemData[];
  focusedItemIndex: { number: number; element: FocusElement };
  addItem: (item: ItemData, index?: number) => void;
  deleteItem: (index: number) => void;
  setItem: (item: ItemData, index: number) => void;
  setFocusedItemIndex: (index?: number, element?: FocusElement) => void;
} {
  const [state, dispatch] = useApp();

  function addItem(item: ItemData, index?: number) {
    dispatch({ type: 'addItem', item, index });
  }

  function deleteItem(index: number) {
    dispatch({ type: 'deleteItem', index });
  }

  function setItem(item: ItemData, index: number) {
    dispatch({ type: 'setItem', item, index });
  }

  function setFocusedItemIndex(index?: number, element?: FocusElement) {
    dispatch({
      type: 'setFocusedItemIndex',
      index,
      element,
    });
  }

  return {
    items: state.items,
    focusedItemIndex: state.focusedItemIndex,
    addItem,
    deleteItem,
    setItem,
    setFocusedItemIndex,
  };
}

export default AppContext;
