import React, {
  createContext,
  RefObject,
  useContext,
  useEffect,
  useMemo,
  useReducer,
  useRef,
} from 'react';
import H5AudioPlayer from 'react-h5-audio-player';
import config from '../../../config/config';

// emptyEvent is passed to functions that require passing an event.
const emptyEvent = {
  stopPropagation: () => {},
} as React.SyntheticEvent;

type AppState = {
  ref: RefObject<H5AudioPlayer> | null;

  audioUri: string;
  isPlaying: boolean;
  duration: number;
};

const initialState: AppState = {
  ref: null,
  audioUri: '',
  isPlaying: false,
  duration: 0,
};

type Action =
  | {
      type: 'setAudioUri';
      audioUri: string;
    }
  | {
      type: 'setIsPlaying';
      isPlaying: boolean;
    }
  | {
      type: 'setRef';
      ref: RefObject<H5AudioPlayer>;
    }
  | {
      type: 'setDuration';
      duration: number;
    };

type Dispatch = (action: Action) => void;

const AudioPlayerContext = createContext<
  { state: AppState; dispatch: Dispatch } | undefined
>(undefined);

function appReducer(state: AppState, action: Action): AppState {
  switch (action.type) {
    case 'setRef': {
      return { ...state, ref: action.ref };
    }

    case 'setAudioUri': {
      // Also set isPlaying true the player plays it when changed
      return { ...state, audioUri: action.audioUri, isPlaying: true };
    }

    case 'setIsPlaying': {
      return { ...state, isPlaying: action.isPlaying };
    }

    case 'setDuration': {
      return { ...state, duration: action.duration };
    }

    default: {
      throw new Error(`Unhandled action: ${action}`);
    }
  }
}

type AudioProviderProps = {
  children: React.ReactNode;
};

export function AudioPlayerProvider({ children }: AudioProviderProps) {
  const [state, dispatch] = useReducer(appReducer, initialState);
  const ref = useRef<H5AudioPlayer>(null);

  const value = useMemo(() => {
    return { state, dispatch };
  }, [state, dispatch]);

  // Initialize ref
  useEffect(() => {
    dispatch({ type: 'setRef', ref });
  }, []);

  // Always jump to start on play
  useEffect(() => {
    if (state.isPlaying) {
      state.ref?.current?.setJumpTime(0);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [state.isPlaying]);

  let audioEndpoint = '';
  if (state.audioUri !== '') {
    audioEndpoint = `${
      config.API_URL
    }/api/v1/app/sample/load/${encodeURIComponent(state.audioUri)}`;
  }

  return (
    <AudioPlayerContext.Provider value={value}>
      <H5AudioPlayer
        ref={ref}
        style={{ display: 'none' }}
        src={audioEndpoint}
        onEnded={() => dispatch({ type: 'setIsPlaying', isPlaying: false })}
        onLoadedMetaData={() => {
          dispatch({
            type: 'setDuration',
            duration: ref.current?.audio?.current?.duration || 0,
          });
        }}
      />
      {children}
    </AudioPlayerContext.Provider>
  );
}

function useAudioPlayerContext(): [AppState, Dispatch] {
  const context = useContext(AudioPlayerContext);
  if (context === undefined) {
    throw new Error('useApp must be used within a AuthContext');
  }
  return [context.state, context.dispatch];
}

export function useAudioPlayer(): {
  audioUri: string;
  setAudioUri: (audioUri: string) => void;

  isPlaying: boolean;
  setIsPlaying: (isPlaying: boolean) => void;

  duration: number;
} {
  const [state, dispatch] = useAudioPlayerContext();

  function setAudioUri(audioUri: string) {
    dispatch({ type: 'setAudioUri', audioUri });
  }

  function setIsPlaying(isPlaying: boolean) {
    state.ref?.current?.togglePlay(emptyEvent);
    dispatch({ type: 'setIsPlaying', isPlaying });
  }

  return {
    audioUri: state.audioUri,
    setAudioUri,
    isPlaying: state.isPlaying,
    setIsPlaying,
    duration: state.duration,
  };
}

export default AudioPlayerContext;
