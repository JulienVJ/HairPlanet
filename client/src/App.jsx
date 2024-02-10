import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './routes/Home';
import Test from './routes/Test';
import Shop from './routes/Shop';
import FormHairdresser from './routes/FormHairdresser';
import Users from './routes/AdminHomePage';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/test" element={<Test />} />
        <Route path="/shop/:name" element={<Shop />} />
        <Route path="/createHairdresser" element={<FormHairdresser />} />
        <Route path="/users" element={<Users />} />
        {/* 
            <Route path="/contact" element={<Contact />} />
            <Route element={<NOT FOUND />} /> This is the 404 route 
        */}
      </Routes>
    </Router>
  );
}

export default App;
