import "./App.css";
import Analyze from "./components/Analyze";
import NavBar from "./components/NavBar";
import { FileProvider } from "./contexts/FileContext";

const App = () => {
  return (
    <div id="app">
      <FileProvider>
        <NavBar />
        <Analyze />
      </FileProvider>
    </div>
  );
};

export default App;
