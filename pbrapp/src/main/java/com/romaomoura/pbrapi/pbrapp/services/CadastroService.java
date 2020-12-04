package com.romaomoura.pbrapi.pbrapp.services;

import java.util.List;

import com.romaomoura.pbrapi.pbrapp.domain.Cadastro;
import com.romaomoura.pbrapi.pbrapp.repositories.CadastroRep;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CadastroService {
    
    @Autowired
    private CadastroRep cadastroRep;

    public Cadastro save(Cadastro cadastro) {
        return cadastroRep.save(cadastro);
    }


    public List<Cadastro> list() {
        return cadastroRep.findAll();
    }
}
