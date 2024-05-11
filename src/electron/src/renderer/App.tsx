import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import Main from './routes/Main/Main';
import { AppProvider } from './components/AppProvider/AppProvider';

export default function App() {
  return (
    <AppProvider>
      <Router>
        <Routes>
          <Route path="/" element={<Main />} />
        </Routes>
      </Router>
    </AppProvider>
  );
}
