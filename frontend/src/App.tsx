import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Body from "./components/Body";
import Navbar from "./components/Navbar";
import Appoint from "./components/Appoint";

export default function App() {

  return (
 
    <Router>
 
      <div>
 
        <Navbar />
        
        <Switch>
 
          <Route exact path="/" component={Body} />
          <Route exact path="/list" component={Appoint} />

        </Switch>
 
      </div>
 
    </Router>
 
  );
 
 }