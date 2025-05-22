// Controle de abas
const tabs = document.querySelectorAll('nav .tab');
const sections = document.querySelectorAll('.section');

tabs.forEach(tab => {
  tab.addEventListener('click', () => {
    tabs.forEach(t => t.classList.remove('active'));
    tab.classList.add('active');
    const target = tab.getAttribute('data-tab');
    sections.forEach(sec => {
      sec.id === target ? sec.classList.add('active') : sec.classList.remove('active');
    });
  });
});

// Função para cadastrar produto
document.getElementById('formCadastrar').addEventListener('submit', async e => {
  e.preventDefault();
  const produto = {
    nome: document.getElementById('nome').value,
    preco: parseFloat(document.getElementById('preco').value),
    quantidade: parseInt(document.getElementById('quantidade').value)
  };

  try {
    const res = await fetch('http://localhost:8080/produtos', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(produto)
    });

    const msg = document.getElementById('mensagemCadastrar');

    if (res.ok) {
      msg.textContent = 'Produto cadastrado com sucesso!';
      e.target.reset();
    } else {
      const erro = await res.json();
      msg.textContent = 'Erro: ' + (erro.erro || 'Erro desconhecido');
    }
  } catch (error) {
    document.getElementById('mensagemCadastrar').textContent = 'Erro na conexão.';
  }
});

// Função para listar produtos
async function listarProdutos() {
  const tabela = document.getElementById('tabelaProdutos');
  tabela.innerHTML = 'Carregando...';

  try {
    const res = await fetch('http://localhost:8080/produtos');
    if (!res.ok) throw new Error('Falha ao buscar produtos');
    const produtos = await res.json();

    if (produtos.length === 0) {
      tabela.innerHTML = '<p>Nenhum produto cadastrado.</p>';
      return;
    }

    let html = '<table><thead><tr><th>ID</th><th>Nome</th><th>Preço</th><th>Quantidade</th></tr></thead><tbody>';
    produtos.forEach(p => {
      html += `<tr><td>${p.id}</td><td>${p.nome}</td><td>R$ ${p.preco.toFixed(2)}</td><td>${p.quantidade}</td></tr>`;
    });
    html += '</tbody></table>';
    tabela.innerHTML = html;
  } catch (error) {
    tabela.innerHTML = '<p>Erro ao carregar produtos.</p>';
  }
}

// Função para atualizar produto
document.getElementById('formAtualizar').addEventListener('submit', async e => {
  e.preventDefault();
  const id = parseInt(document.getElementById('idAtualizar').value);
  const novosDados = {};

  const nome = document.getElementById('nomeAtualizar').value;
  if (nome) novosDados.nome = nome;
  const preco = document.getElementById('precoAtualizar').value;
  if (preco) novosDados.preco = parseFloat(preco);
  const quantidade = document.getElementById('quantidadeAtualizar').value;
  if (quantidade) novosDados.quantidade = parseInt(quantidade);

  if (Object.keys(novosDados).length === 0) {
    document.getElementById('mensagemAtualizar').textContent = 'Preencha ao menos um campo para atualizar.';
    return;
  }

  try {
    const res = await fetch(`http://localhost:8080/produtos/${id}`, {
      method: 'PUT',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(novosDados)
    });
    const msg = document.getElementById('mensagemAtualizar');

    if (res.ok) {
      msg.textContent = 'Produto atualizado com sucesso!';
      e.target.reset();
    } else {
      const erro = await res.json();
      msg.textContent = 'Erro: ' + (erro.erro || 'Erro desconhecido');
    }
  } catch {
    document.getElementById('mensagemAtualizar').textContent = 'Erro na conexão.';
  }
});

// Função para deletar produto
async function deletarProduto() {
  const id = parseInt(document.getElementById('idDeletar').value);
  if (!id) {
    document.getElementById('mensagemDeletar').textContent = 'Informe um ID válido.';
    return;
  }

  if (!confirm(`Confirma a remoção do produto com ID ${id}?`)) {
    return;
  }

  try {
    const res = await fetch(`http://localhost:8080/produtos/${id}`, {
      method: 'DELETE'
    });
    const msg = document.getElementById('mensagemDeletar');

    if (res.ok) {
      msg.textContent = 'Produto deletado com sucesso!';
      document.getElementById('idDeletar').value = '';
    } else {
      const erro = await res.json();
      msg.textContent = 'Erro: ' + (erro.erro || 'Erro desconhecido');
    }
  } catch {
    document.getElementById('mensagemDeletar').textContent = 'Erro na conexão.';
  }
}
