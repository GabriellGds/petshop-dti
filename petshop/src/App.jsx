import React, { useState } from 'react'
import './App.css' 
import caninoFelix from './assets/canino-felix.jpg'
import chowChawgas from './assets/chowChawgas.jpg'
import rex from './assets/rex.jpg'

function App() {
  const [date, setDate] = useState('')
  const [smallDogs, setSmallDogs] = useState(0)
  const [bigDogs, setBigDogs] = useState(0)
  const [result, setResult] = useState('')
  const [error, setError] = useState('')

  const getBestPetshop = async () => {
    try {
      const formattedDate = date.split('-').join('/');
      const response = await fetch(`http://localhost:8000/api/petshops?date=${formattedDate}&smallDogs=${smallDogs}&bigDogs=${bigDogs}`)
      const data = await response.json()
      setResult(data)
      setError('')
    } catch (err) {
      console.error(err)
      setError('Por favor forneça todos os campos.')
    }
  }

  return (
    <div className="app">
      <h1>Encontre o melhor petshop</h1>
      <div className="petshops">
        <div className="petshop">
          <div className="image-container">
            <img src={caninoFelix} alt="Meu Canino Feliz" />
          </div>
          <h3>Meu Canino Feliz</h3>
          <p>Em dias de semana o banho para cães pequenos custa R$20,00 e o banho para cães grandes custa R$40,00. Durante os finais de semana o preço dos banhos é aumentado em 20%.</p>
        </div>
        <div className="petshop">
          <div className="image-container">
            <img src={chowChawgas} alt="ChowChawgas" />
          </div>
          <h3>Vai Rex</h3>
          <p>O preço do banho para dias úteis em cães pequenos é R$15,00 e em cães grandes é R$50,00. Durante os finais de semana o preço para cães pequenos é R$20,00 e para os grandes é R$55,00.</p>
        </div>
        <div className="petshop">
          <div className="image-container">
            <img src={rex} alt="Vai rex" />
          </div>
          <h3>ChowChawgas</h3>
          <p>O preço do banho é o mesmo em todos os dias da semana. Para cães pequenos custa R$30 e para cães grandes R$45,00.</p>
        </div>
      </div>
      <div className="input-group">
        <label>Data:</label>
        <input type="date" value={date} onChange={e => setDate(e.target.value)} />
      </div>
      <div className="input-group">
        <label>Número de cães pequenos:</label>
        <input type="number" value={smallDogs} onChange={e => setSmallDogs(e.target.value)} />
      </div>
      <div className="input-group">
        <label>Número de cães grandes:</label>
        <input type="number" value={bigDogs} onChange={e => setBigDogs(e.target.value)} />
      </div>
      <button onClick={getBestPetshop}>Buscar</button>
      {error ? (
        <div className="error">
          <h2>Erro</h2>
          <p>{error}</p>
        </div>
      ) : result && (
        <div className="result">
          <h2>Melhor Petshop</h2>
          <p>Petshop: {result.petshop}</p>
          <p>Preço Total: {result.totalPrice}R$</p>
        </div>
      )}
    </div>
  );
}

export default App;
