package com.romaomoura.pbrapi.pbrapp.repositories;

import com.romaomoura.pbrapi.pbrapp.domain.Cadastro;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CadastroRep extends JpaRepository<Cadastro, Long> {
    
}
