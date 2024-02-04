import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './routes/Home';
import Test from './routes/Test';
import Shop from './routes/Shop';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/test" element={<Test />} />
        <Route path="/shop/:name" element={<Shop />} />
        {/* 
            <Route path="/contact" element={<Contact />} />
            <Route element={<NOT FOUND />} /> This is the 404 route 
        */}
      </Routes>
    </Router>
  );
}

export default App;
