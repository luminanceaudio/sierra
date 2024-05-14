import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import Main from './routes/Main/Main';
import { AudioPlayerProvider } from './components/AudioPlayerProvider/AudioPlayerProvider';
import { PersistQueryClientProvider } from '../api/reactquery';
import Sources from './routes/Sources/Sources';

export default function App() {
  return (
    <PersistQueryClientProvider>
      <AudioPlayerProvider>
        <Router>
          <Routes>
            <Route path="/" element={<Main />} />
            <Route path="/sources" element={<Sources />} />
          </Routes>
        </Router>
      </AudioPlayerProvider>
    </PersistQueryClientProvider>
  );
}
