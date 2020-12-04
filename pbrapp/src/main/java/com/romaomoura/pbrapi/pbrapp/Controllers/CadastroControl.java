package com.romaomoura.pbrapi.pbrapp.Controllers;

import com.romaomoura.pbrapi.pbrapp.domain.Cadastro;
import com.romaomoura.pbrapi.pbrapp.repositories.CadastroRep;
import com.romaomoura.pbrapi.pbrapp.services.CadastroService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("cadastrar")
public class CadastroControl {

    @Autowired
    private CadastroService cadastroServ;

    @GetMapping("/") //redireciona para pagina formulario
    public String cadastro() {
        return "/novo_cadastro";
    }

    @PostMapping("/salvar") //salva um novo cadastro
    public String salvarcadastro(@ModelAttribute("cadastro") Cadastro cadastro) {

        cadastroServ.save(cadastro);
        return "/index";
    }

    @GetMapping("/listar") //listar todos os cadastros
    public String list(ModelMap model) {
        model.addAttribute("cadastros", cadastroServ.list());
        return "/listar-todos";
    }
    
}
