import { Provider } from 'mobx-react';
import objectStore from './stores/ObjectsStore';
import MainPage from './components/MainPage';

const stores = { objectStore };

function App() {
  return (
    <Provider { ...stores }>
      <MainPage />
    </Provider>
  );
}

export default App;
