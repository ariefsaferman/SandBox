const data = [];

document.getElementById("form_id").addEventListener("submit", function (event) {
  event.preventDefault();
  let form = event.target.elements;
  let name = form.name.value;
  let email = form.email.value;
  let phone = form.phone.value;
  let message = form.message.value;
  const contact = {
    name: name,
    email: email,
    phone: phone,
    message: message,
  };

  data.push(contact);
  fillTable();
});

const fillTable = () => {
  const table = document.getElementById("table_id");
  table.innerHTML = "";
  data.forEach(function (value, index) {
    table.innerHTML += `<tr>
    <td>${index + 1}</td>
    <td>${value.name}</td>
    <td>${value.email}</td>
    <td>${value.message}</td>
    <td><button class="btn btn-danger" onclick="deleteItem(${index})">Delete </button></td>
</tr>`;
  });
};

const deleteItem = (id) => {
  data.splice(id);
  fillTable();
};
