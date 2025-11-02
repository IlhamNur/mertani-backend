const API_URL = "http://localhost:8082/sensors";
const tableBody = document.querySelector("#sensorTable tbody");
const modal = new bootstrap.Modal(document.getElementById("sensorModal"));

async function fetchSensors() {
  const res = await fetch(API_URL);
  const data = await res.json();
  tableBody.innerHTML = "";
  data.forEach((s) => {
    const row = `<tr>
      <td>${s.id}</td>
      <td>${s.deviceId}</td>
      <td>${s.type}</td>
      <td>${s.value}</td>
      <td>${s.unit}</td>
      <td>
        <button class="btn btn-sm btn-warning" onclick="editSensor(${s.id})">Edit</button>
        <button class="btn btn-sm btn-danger" onclick="deleteSensor(${s.id})">Hapus</button>
      </td>
    </tr>`;
    tableBody.insertAdjacentHTML("beforeend", row);
  });
}

async function saveSensor() {
  const id = document.getElementById("sensorId").value;
  const payload = {
    deviceId: parseInt(document.getElementById("sensorDeviceId").value),
    type: document.getElementById("sensorType").value,
    value: parseFloat(document.getElementById("sensorValue").value),
    unit: document.getElementById("sensorUnit").value,
  };
  const method = id ? "PUT" : "POST";
  const url = id ? `${API_URL}/${id}` : API_URL;

  await fetch(url, {
    method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  modal.hide();
  document.querySelector("button[data-bs-target='#sensorModal']").focus();
  fetchSensors();
}

async function editSensor(id) {
  const res = await fetch(`${API_URL}/${id}`);
  const s = await res.json();
  document.getElementById("sensorId").value = s.id;
  document.getElementById("sensorDeviceId").value = s.deviceId;
  document.getElementById("sensorType").value = s.type;
  document.getElementById("sensorValue").value = s.value;
  document.getElementById("sensorUnit").value = s.unit;
  modal.show();
}

async function deleteSensor(id) {
  if (!confirm("Yakin hapus sensor ini?")) return;
  await fetch(`${API_URL}/${id}`, { method: "DELETE" });
  fetchSensors();
}

document.getElementById("saveSensorBtn").addEventListener("click", saveSensor);
document.addEventListener("DOMContentLoaded", fetchSensors);
