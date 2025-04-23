import ReactDOM from "react-dom/client";
import { App } from "./App";


const rootElement = document.getElementById("root")!;
const root = ReactDOM.createRoot(rootElement);

const app = (
    <div className="root-child flex flex-wrap">
        <App />
    </div>
)

root.render(app);
