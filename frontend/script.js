let produtosCache = [];

function mostrarSecao(secao) {
  document.querySelectorAll(".secao").forEach(s => s.style.display = "none");
  document.getElementById(secao).style.display = "block";

  if (secao === "listar") carregarProdutos();
  if (secao === "atualizar") carregarSelectAtualizacao();
}

function cadastrarProduto() {
  const nome = document.getElementById("nome").value;
  const preco = parseFloat(document.getElementById("preco").value);
  const quantidade = parseInt(document.getElementById("quantidade").value);

  fetch("http://localhost:8080/produtos", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nome, preco, quantidade })
  })
    .then(res => res.ok ? "Produto cadastrado com sucesso!" : "Erro ao cadastrar produto")
    .then(msg => {
      document.getElementById("mensagemCadastro").textContent = msg;
      document.getElementById("nome").value = "";
      document.getElementById("preco").value = "";
      document.getElementById("quantidade").value = "";
    });
}

function carregarProdutos() {
  fetch("http://localhost:8080/produtos")
    .then(res => res.json())
    .then(produtos => {
      const tbody = document.querySelector("#tabelaProdutos tbody");
      tbody.innerHTML = "";
      produtos.forEach(p => {
        tbody.innerHTML += `<tr><td>${p.id}</td><td>${p.nome}</td><td>${p.preco}</td><td>${p.quantidade}</td></tr>`;
      });
    });
}

function carregarSelectAtualizacao() {
  fetch("http://localhost:8080/produtos")
    .then(res => res.json())
    .then(produtos => {
      const select = document.getElementById("produtoSelect");
      select.innerHTML = '<option value="">Selecione um produto</option>';
      produtosCache = produtos;

      produtos.forEach(p => {
        const option = document.createElement("option");
        option.value = p.id;
        option.text = `${p.nome} - R$${p.preco} - Qtd: ${p.quantidade}`;
        select.appendChild(option);
      });
    });
}

document.getElementById("produtoSelect").addEventListener("change", () => {
  const id = document.getElementById("produtoSelect").value;
  const produto = produtosCache.find(p => p.id == id);
  if (produto) {
    document.getElementById("novoNome").value = produto.nome;
    document.getElementById("novoPreco").value = produto.preco;
    document.getElementById("novaQtd").value = produto.quantidade;
  }
});

function atualizarProduto() {
  const id = document.getElementById("produtoSelect").value;
  const nome = document.getElementById("novoNome").value;
  const preco = parseFloat(document.getElementById("novoPreco").value);
  const quantidade = parseInt(document.getElementById("novaQtd").value);

  if (!id) return alert("Selecione um produto");

  fetch(`http://localhost:8080/produtos/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nome, preco, quantidade })
  })
    .then(res => res.ok ? "Produto atualizado com sucesso!" : "Erro ao atualizar produto")
    .then(msg => {
      document.getElementById("mensagemAtualizar").textContent = msg;
      carregarSelectAtualizacao();
    });
}

function deletarProduto() {
  const id = document.getElementById("idDeletar").value;
  if (!id) return alert("Informe o ID do produto");

  if (confirm("Tem certeza que deseja deletar este produto?")) {
    fetch(`http://localhost:8080/produtos/${id}`, {
      method: "DELETE"
    })
      .then(res => res.ok ? "Produto deletado com sucesso!" : "Erro ao deletar")
      .then(msg => {
        document.getElementById("mensagemDeletar").textContent = msg;
        document.getElementById("idDeletar").value = "";
        carregarProdutos();
      });
  }
}
