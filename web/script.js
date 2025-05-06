
const apiurl = 'http://127.0.0.1:3000';

async function loadPackSizes() {
    const response = await fetch(apiurl+'/api/pack-sizes');
    const data = await response.json();
    document.getElementById('packSizes').value = data.packSizes.join('\n');
}

async function updatePackSizes() {
    const packSizesText = document.getElementById('packSizes').value;
    const packSizes = packSizesText.split('\n').map(Number).filter(n => n > 0);
    const response = await fetch(apiurl+'/api/pack-sizes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ packSizes })
    });
    const result = await response.json();
    alert(result.message || result.error);
}

async function calculatePacks() {
    const orderAmount = Number(document.getElementById('orderAmount').value);
    const response = await fetch(apiurl+'/api/calculate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ orderAmount })
    });
    const result = await response.json();

    if (result.error) {
        alert(result.error);
        return;
    }

    const tableBody = document.getElementById('resultTable');
    tableBody.innerHTML = '';
    for (const [pack, quantity] of Object.entries(result.packs)) {
        const row = document.createElement('tr');
        row.innerHTML = `<td>${pack}</td><td>${quantity}</td>`;
        tableBody.appendChild(row);
    }
    document.getElementById('totalItems').textContent = result.totalItems;
}

// Load pack sizes on page load
loadPackSizes();