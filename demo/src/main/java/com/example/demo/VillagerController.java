package com.example.demo;


import lombok.extern.slf4j.Slf4j;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@Slf4j
public class VillagerController {


    @PostMapping(value = "/villager", consumes = {MediaType.APPLICATION_JSON_VALUE} )
    public Villager villager (@RequestBody Villager villager) {
        log.info("Welcome {} !!" , villager);
        log.info("Welcome {} !!" , villager.getName());
        return villager;
    }
}
