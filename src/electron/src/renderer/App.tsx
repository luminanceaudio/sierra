import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import Main from './routes/Main/Main';
import { AppProvider } from './components/AppProvider/AppProvider';
import { PersistQueryClientProvider } from '../api/reactquery';

export default function App() {
  return (
    <AppProvider>
      <PersistQueryClientProvider>
        <Router>
          <Routes>
            <Route path="/" element={<Main />} />
          </Routes>
        </Router>
      </PersistQueryClientProvider>
    </AppProvider>
  );
}
