function spin() {
  fetch('/spin')
    .then((response) => response.text())
    .then((data) => {
      document.getElementById('result').innerText = data;
      checkBalance(); // Update balance after spin
    })
    .catch((error) => console.error('Error during spin:', error));
}

function checkBalance() {
  fetch('/balance')
    .then((response) => response.json())
    .then((data) => {
      document.getElementById('balance').innerText = 'Balance: $' + data.balance.toFixed(2);
    })
    .catch((error) => console.error('Error fetching balance:', error));
}

// Initially load balance
document.addEventListener('DOMContentLoaded', checkBalance);
