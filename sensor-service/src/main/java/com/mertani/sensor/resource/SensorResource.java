package com.mertani.sensor.resource;

import jakarta.ws.rs.core.MediaType;
import jakarta.inject.Inject;
import jakarta.transaction.Transactional;
import jakarta.ws.rs.*;
import java.util.List;
import com.mertani.sensor.entity.Sensor;
import com.mertani.sensor.repository.SensorRepository;

@Path("/sensors")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class SensorResource {

    @Inject
    SensorRepository repo;

    @GET
    public List<Sensor> getAll() {
        return repo.listAll();
    }

    @GET
    @Path("/{id}")
    public Sensor get(@PathParam("id") Long id) {
        return repo.findById(id);
    }

    @POST
    @Transactional
    public Sensor create(Sensor s) {
        s.createdAt = java.time.LocalDateTime.now();
        s.updatedAt = java.time.LocalDateTime.now();
        repo.persist(s);
        return s;
    }

    @PUT
    @Path("/{id}")
    @Transactional
    public Sensor update(@PathParam("id") Long id, Sensor s) {
        Sensor exist = repo.findById(id);
        if (exist == null)
            throw new NotFoundException();
        exist.type = s.type;
        exist.value = s.value;
        exist.unit = s.unit;
        exist.deviceId = s.deviceId;
        exist.updatedAt = java.time.LocalDateTime.now();
        repo.persist(exist);
        return exist;
    }

    @DELETE
    @Path("/{id}")
    @Transactional
    public void delete(@PathParam("id") Long id) {
        repo.deleteById(id);
    }
}
