const API_URL = "http://localhost:8080/devices";
const tableBody = document.querySelector("#deviceTable tbody");
const modal = new bootstrap.Modal(document.getElementById("deviceModal"));

async function fetchDevices() {
  const res = await fetch(API_URL);
  const data = await res.json();
  tableBody.innerHTML = "";
  data.forEach((d) => {
    const row = `<tr>
      <td>${d.id}</td>
      <td>${d.name}</td>
      <td>${d.location}</td>
      <td>${d.status ? "✅" : "❌"}</td>
      <td>
        <button class="btn btn-sm btn-warning" onclick="editDevice(${
          d.id
        })">Edit</button>
        <button class="btn btn-sm btn-danger" onclick="deleteDevice(${
          d.id
        })">Hapus</button>
      </td>
    </tr>`;
    tableBody.insertAdjacentHTML("beforeend", row);
  });
}

async function saveDevice() {
  const id = document.getElementById("deviceId").value;
  const payload = {
    name: document.getElementById("deviceName").value,
    location: document.getElementById("deviceLocation").value,
    status: document.getElementById("deviceStatus").checked,
  };
  const method = id ? "PUT" : "POST";
  const url = id ? `${API_URL}/${id}` : API_URL;

  await fetch(url, {
    method,
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  modal.hide();
  document.querySelector("button[data-bs-target='#deviceModal']").focus();
  fetchDevices();
}

async function editDevice(id) {
  const res = await fetch(`${API_URL}/${id}`);
  const d = await res.json();
  document.getElementById("deviceId").value = d.id;
  document.getElementById("deviceName").value = d.name;
  document.getElementById("deviceLocation").value = d.location;
  document.getElementById("deviceStatus").checked = d.status;
  modal.show();
}

async function deleteDevice(id) {
  if (!confirm("Yakin hapus device ini?")) return;
  await fetch(`${API_URL}/${id}`, { method: "DELETE" });
  fetchDevices();
}

document.getElementById("saveDeviceBtn").addEventListener("click", saveDevice);
document.addEventListener("DOMContentLoaded", fetchDevices);
