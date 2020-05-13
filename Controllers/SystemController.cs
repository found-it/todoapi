using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Net;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace TodoApi.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    [Produces("application/json")]
    public class SystemController : ControllerBase
    {
        private readonly ILogger<SystemController> _logger;

        public SystemController(ILogger<SystemController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public string Get()
        {
            String hostname = Dns.GetHostName();
            IPHostEntry ip = Dns.GetHostEntry(hostname);
            IPAddress[] address = ip.AddressList;

            // String s = String.Format("{\"hostname\": \"{0}\", \"ip\": [", hostname.ToString());
            String s = "{\"hostname\":\"" + hostname + "\", \"ip\":[";

            for (int i = 0; i < address.Length; ++i)
            {
                s += "\"";
                s += address[i].ToString();
                s += "\"";
                if (i < address.Length - 1)
                {
                    s += ",";
                }
            }
            s += "]}";



            return s;
        }
    }
}