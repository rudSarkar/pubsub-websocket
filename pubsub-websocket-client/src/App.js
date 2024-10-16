import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Order from "./routes/Order";
import Bill from "./routes/Bill";
import './index.css';
import Home from './routes/Home';

function App() {
  return (
    <Router>
      <div className="container mx-auto">
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path="/order" element={<Order />} />
          <Route path="/bill" element={<Bill />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
